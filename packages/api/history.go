/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package api

import (
	"encoding/json"
	"net/http"

	"github.com/IBAX-io/go-ibax/packages/consts"
	"github.com/IBAX-io/go-ibax/packages/model"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const rollbackHistoryLimit = 100

type historyResult struct {
	List []map[string]string `json:"list"`
}

func getHistoryHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	logger := getLogger(r)
	client := getClient(r)

	table := client.Prefix() + "_" + params["name"]
	rollbackTx := &model.RollbackTx{}
	txs, err := rollbackTx.GetRollbackTxsByTableIDAndTableName(params["id"], table, rollbackHistoryLimit)
	if err != nil {
		logger.WithFields(log.Fields{"type": consts.DBError, "error": err}).Error("rollback history")
		errorResponse(w, err)
}