package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"path"
	"strconv"

	"../wallet"
)

// WalletServer -> Type Definition.
type WalletServer struct {
	port    uint16
	gateway string
}

// NewWalletServer -> Create a new wallet server.
func NewWalletServer(port uint16, gateway string) *WalletServer {
	return &WalletServer{port, gateway}
}

// Port -> Return the Port of the wallet server.
func (ws *WalletServer) Port() uint16 {
	return ws.port
}

// Gateway -> Return the Gateway of the wallet server.
func (ws *WalletServer) Gateway() string {
	return ws.gateway
}

// Index ->
func (ws *WalletServer) Index(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		t, _ := template.ParseFiles(path.Join("walletServer/templates", "index.html"))
		t.Execute(w, "")
	default:
		log.Printf("ERROR: Invalid HTTP Method")
	}
}

// Wallet ->
func (ws *WalletServer) Wallet(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		w.Header().Add("Content-Type", "application/json")
		myWallet := wallet.NewWallet()
		m, _ := myWallet.MarshalJSON()
		io.WriteString(w, string(m[:]))
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Invalid HTTP Method")
	}
}

// Run -> Run the wallet server.
func (ws *WalletServer) Run() {
	http.HandleFunc("/", ws.Index)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.Port())), nil))
}
