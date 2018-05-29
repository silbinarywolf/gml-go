#DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" ) " && pwd )"
protoc --gofast_out=. map_entity.proto map.proto
#protoc --gofast_out=. map.proto
