package main

import (
	"fmt"
	"net"
)

type Cliente struct {
	conectado bool
	conexao   net.Conn
}

var listener net.Listener
var clientes []*Cliente

const tamanhoBufferDasMensagens = 4096

func IniciarServidorSocket() {
	serverAux, err := net.Listen("tcp", ":60000")
	if err != nil {
		panic(fmt.Sprintf("listener error %s", err))
	}

	fmt.Println("servidor iniciado na porta 60000")
	listener = serverAux
}

func AceitarConexao(cliente chan *Cliente) {
	go func() {
		for {
			fmt.Println("aguardando nova conexão cliente")
			con, err := listener.Accept()
			if err != nil {
				fmt.Println("Falha ao aceitar conexão do cliente")
				continue
			}
			fmt.Printf("nova conexão cliente de: %s\n", con.RemoteAddr().String())
			novoCliente := Cliente{conexao: con}
			clientes = append(clientes, &novoCliente)
			cliente <- &novoCliente
		}
	}()
}

func (c *Cliente) encerrarConexao() {
	c.conectado = false
	_ = c.conexao.Close()
	for index, cli := range clientes {
		if cli == c {
			clientes[index].conexao = nil
			fmt.Printf("conexão cliente %d encerrrada", index)
		}
	}
}

func (c *Cliente) ReceberMensagens() {
	defer c.encerrarConexao()
	mensagemBuffer := make([]byte, tamanhoBufferDasMensagens)
	for {
		tamanhoMensagemRecebida, err := c.conexao.Read(mensagemBuffer)
		if err != nil {
			break
		}

		fmt.Println(string(mensagemBuffer))
		_, _ = c.conexao.Write(mensagemBuffer[:tamanhoMensagemRecebida])
	}
}

func main() {
	IniciarServidorSocket()
	var clienteChannel chan *Cliente
	clienteChannel = make(chan *Cliente)
	AceitarConexao(clienteChannel)
	for {
		cliente := <-clienteChannel
		go cliente.ReceberMensagens()
	}
}
