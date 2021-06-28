package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers blockchaintest-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	// this line is used by starport scaffolding # 1
	r.HandleFunc("/blockchaintest/whois", buyNameHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/blockchaintest/whois", listWhoisHandler(cliCtx, "blockchaintest")).Methods("GET")
	r.HandleFunc("/blockchaintest/whois/{key}", getWhoisHandler(cliCtx, "blockchaintest")).Methods("GET")
	r.HandleFunc("/blockchaintest/whois/{key}/resolve", resolveNameHandler(cliCtx, "blockchaintest")).Methods("GET")
	r.HandleFunc("/blockchaintest/whois", setNameHandler(cliCtx)).Methods("PUT")
	r.HandleFunc("/blockchaintest/whois", deleteNameHandler(cliCtx)).Methods("DELETE")

}
