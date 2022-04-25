package main

import (
    "fmt"
    "net/http"
    "time"

    "github.com/iotaledger/goshimmer/client"
    walletseed "github.com/iotaledger/goshimmer/client/wallet/packages/seed"
    "github.com/iotaledger/goshimmer/packages/ledgerstate"
    "github.com/iotaledger/goshimmer/packages/mana"
)

func buildTransaction() (tx *ledgerstate.Transaction, err error) {
    // node to pledge access mana.
    accessManaPledgeIDBase58 := "FZ6xmPZXRs2M8z9m9ETTQok4PCga4X8FRHwQE6uYm4rV"
    accessManaPledgeID, err := mana.IDFromStr(accessManaPledgeIDBase58)
    if err != nil {
        return
    }

    // node to pledge consensus mana.
    consensusManaPledgeIDBase58 := "FZ6xmPZXRs2M8z9m9ETTQok4PCga4X8FRHwQE6uYm4rV"
    consensusManaPledgeID, err := mana.IDFromStr(consensusManaPledgeIDBase58)
    if err != nil {
        return
    }
     
        /**
        N.B to pledge mana to the node issuing the transaction, use empty pledgeIDs.
        emptyID := identity.ID{}
        accessManaPledgeID, consensusManaPledgeID := emptyID, emptyID
        **/      

    // destination address.
    destAddressBase58 := "1F5UHGuVcyDBMRUvg8QbBLcNLJKWQE9pBLyUq7wBxzxtz"
    destAddress, err := ledgerstate.AddressFromBase58EncodedString(destAddressBase58)
    if err != nil {
        return
    }

    // output to consume.
    outputIDBase58 := "18UkRSsSNUCwbHU3jaP3HxfyeBhJS3vZcgUoS7S9ah42N"
    out, err := ledgerstate.OutputIDFromBase58(outputIDBase58)
    if err != nil {
        return
    }
    inputs := ledgerstate.NewInputs(ledgerstate.NewUTXOInput(out))

    // UTXO output.
    output := ledgerstate.NewSigLockedColoredOutput(ledgerstate.NewColoredBalances(map[ledgerstate.Color]uint64{
        ledgerstate.ColorIOTA: uint64(1337),
    }), destAddress)
    outputs := ledgerstate.NewOutputs(output)

    // build tx essence.
    txEssence := ledgerstate.NewTransactionEssence(0, time.Now(), accessManaPledgeID, consensusManaPledgeID, inputs, outputs)

    // sign.
    seed := walletseed.NewSeed([]byte("BZx7rpp2FU6etRQL3sWyQAyfmMxRpgBJ8rCgHjNZpucB"))
    kp := seed.KeyPair(0)
    sig := ledgerstate.NewED25519Signature(kp.PublicKey, kp.PrivateKey.Sign(txEssence.Bytes()))
    unlockBlock := ledgerstate.NewSignatureUnlockBlock(sig)

    // build tx.
    tx = ledgerstate.NewTransaction(txEssence, ledgerstate.UnlockBlocks{unlockBlock})
    return
}

// build tx from previous step
tx, err := buildTransaction()
if err != nil {
    return
}
bytes := tx.Bytes()

// print bytes
fmt.Println(string(bytes))