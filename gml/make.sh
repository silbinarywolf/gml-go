#DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" ) " && pwd )"
protoc --gofast_out=. room_instance.proto room.proto
