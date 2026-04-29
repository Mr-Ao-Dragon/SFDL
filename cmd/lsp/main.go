package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"

	sfdl "github.com/lightDproject/SFDL"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		msg, err := readMessage(reader)
		if err != nil {
			if err != io.EOF {
				fmt.Fprintf(os.Stderr, "Read error: %v\n", err)
			}
			return
		}

		var req Request
		if err := json.Unmarshal(msg, &req); err != nil {
			fmt.Fprintf(os.Stderr, "Parse error: %v\n", err)
			continue
		}

		resp := handleRequest(req)
		respBytes, err := json.Marshal(resp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Marshal error: %v\n", err)
			continue
		}
		writeMessage(writer, respBytes)
	}
}

func readMessage(r *bufio.Reader) ([]byte, error) {
	var length int
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			return nil, err
		}
		line = line[:len(line)-1]
		if line == "" {
			break
		}
		if _, err := fmt.Sscanf(line, "Content-Length: %d", &length); err == nil {
			break
		}
	}

	content := make([]byte, length)
	_, err := io.ReadFull(r, content)
	return content, err
}

func writeMessage(w *bufio.Writer, data []byte) {
	header := fmt.Sprintf("Content-Length: %d\r\n\r\n", len(data))
	w.WriteString(header)
	w.Write(data)
	w.Flush()
}

func handleRequest(req Request) Response {
	id := req.ID
	switch req.Method {
	case "initialize":
		return Response{ID: id, Result: InitializeResult{
			Capabilities: ServerCapabilities{
				TextDocumentSync:   1,
				HoverProvider:      true,
				CompletionProvider: true,
			},
		}}
	case "textDocument/didOpen":
		return Response{ID: id}
	case "textDocument/didChange":
		return Response{ID: id}
	case "textDocument/hover":
		return Response{ID: id, Result: Hover{Contents: "SFDL Configuration"}}
	case "textDocument/completion":
		return Response{ID: id, Result: CompletionList{
			Items: func() []CompletionItem {
				types := sfdl.GetRegisteredBlockTypes()
				items := make([]CompletionItem, 0, len(types))
				for _, t := range types {
					items = append(items, CompletionItem{Label: t})
				}
				return items
			}(),
		}}
	default:
		return Response{ID: id}
	}
}

type Request struct {
	ID     any             `json:"id"`
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
}

type Response struct {
	ID     any `json:"id"`
	Result any `json:"result,omitzero"`
	Error  any `json:"error,omitzero"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
}

type ServerCapabilities struct {
	TextDocumentSync   int  `json:"textDocumentSync"`
	HoverProvider      bool `json:"hoverProvider"`
	CompletionProvider bool `json:"completionProvider"`
}

type Hover struct {
	Contents any `json:"contents"`
}

type CompletionList struct {
	Items []CompletionItem `json:"items"`
}

type CompletionItem struct {
	Label string `json:"label"`
}

var _ = sfdl.Config{}
