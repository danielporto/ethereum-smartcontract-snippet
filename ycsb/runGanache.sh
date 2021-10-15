./ycsb deploy --host localhost --port 7545 --key "7d6419430c8f84d5305cb060c301214218338a8aa327ce4a7dc428397a17e930"

# Nonce: 0
# Gas price: 20000000000
# Transaction address: 0x843ef8f573414ae4631cf63ff95a05be90bd10f57c3b300abd5adcd854fae077
# Contract address 0x24C4fbcd6fa5360Cd1AEA2CeD9F1D7a0cdbc1897
# Contract deployed

./ycsb workload --operation ycsb --host 127.0.0.1 --port 8545 --key a3cc08be595b8f4ce8cc76bab3f26eb20acdf82c55397711adcc5977341b9b25 --count 100 --dbkeysize 100 --dbvaluesize 200 --threads 10