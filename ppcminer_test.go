package main

import (
	"testing"
	//"main"
	"bufio"
	"github.com/conformal/btcec"
	"github.com/mably/btcutil"
	"math/rand"
	"os"
	"time"
)

func TestMiner(t *testing.T) {
	initSeelogLogger("/tmp/logs")
	setLogLevels("trace")
	cfg = new(config)
	cfg.DbType = "memdb"
	activeNetParams = &testNet3Params

	db, err := loadBlockDB()
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	defer db.Close()

	// Ensure the database is sync'd and closed on Ctrl+C.
	addInterruptHandler(func() {
		btcdLog.Infof("Gracefully shutting down the database...")
		db.RollbackClose()
	})

	s := server{
		db:        db,
		netParams: activeNetParams.Params,
	}

	bm, err := newBlockManager(&s)
	if err != nil {
		t.Error(err)
		return
	}
	s.blockManager = bm
	s.txMemPool = newTxMemPool(&s)

	bm.Start()

	m := newPPCMiner(&s)

	pk := make([]byte, 32)
	for i := 0; i < 32; i++ {
		pk[i] = byte(rand.Int())
	}
	priv, pub := btcec.PrivKeyFromBytes(btcec.S256(), pk)
	addr, err := btcutil.NewAddressPubKey(pub.SerializeCompressed(), activeNetParams.Params)
	if err != nil {
		t.Error(err)
		return
	}
	m.payToAddress = addr
	m.Start()
	time.Sleep(time.Minute)
	m.Stop()
	time.Sleep(time.Second * 5)

	_ = priv
}
