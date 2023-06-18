import sys
import os

def main(command):
    if command == "test":
        # do testing
        #go test .\contract_modules\uniswap_v2
        print("testing")
        os.system('go test ./contract_modules/uniswap_v2/.')
        os.system('go build ./cmd/bot/.')
        print("running")        
        os.system('bot')

        #foundry test --args="--arg1=value1 --arg2=value2"

        #create fork with foundry
        #generate calldata with go
        #deploy contract with foundry on fork
        #

    elif command == "buildDebugRun":
        print("building")
        os.system('go build -gcflags "all=-N -l" ./cmd/bot/.')
        print("running")        
        os.system('dlv debug bot')

    elif command == "buildDebugRunDataCollector":
        print("building")
        os.system('go build -gcflags "all=-N -l" ./cmd/data_collector/.')
        print("running")        
        os.system('dlv debug data_collector')

    elif command == "build":
        print("building")
        os.system('go build ./cmd/bot/.')

    elif command == "buildRun":
        print("building")
        os.system('go build ./cmd/bot/.')
        print("running")        
        os.system('bot')


if len(sys.argv) == 1:
    print("test / build")
    exit(1)
main(sys.argv[1])

