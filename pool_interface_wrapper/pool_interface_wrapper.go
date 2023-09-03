package pool_interface_wrapper

import (
	"encoding/json"
	"fmt"
	"mev-template-go/pool_interface"
	"mev-template-go/uniswap_v2"
	"mev-template-go/uniswap_v3"
)

type PoolInterfaceWrapper struct {
	pool_interface.PoolInterface
}

func (pw PoolInterfaceWrapper) MarshalJSON() ([]byte, error) {
	var poolJSON struct {
		Type string
		Data pool_interface.PoolInterface
	}

	poolJSON.Type = pw.PoolInterface.GetType()
	poolJSON.Data = pw.PoolInterface

	jsonData, err := json.Marshal(poolJSON)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func (pw *PoolInterfaceWrapper) UnmarshalJSON(b []byte) error {
	var poolJSON struct {
		Type string
		Data json.RawMessage
	}

	err := json.Unmarshal(b, &poolJSON)
	if err != nil {
		return err
	}

	switch poolJSON.Type {
	case "uniswap_v2":
		var poolA uniswap_v2.Pool
		err = json.Unmarshal(poolJSON.Data, &poolA)
		pw.PoolInterface = &poolA
	case "uniswap_v3":
		var poolB uniswap_v3.Pool
		err = json.Unmarshal(poolJSON.Data, &poolB)
		pw.PoolInterface = &poolB
	default:
		err = fmt.Errorf("unknown pool type: %s", poolJSON.Type)
	}

	return err
}
