package serializer_test

import (
	"testing"

	"github.com/bugg123/grpc-intro/pb"
	"github.com/bugg123/grpc-intro/sample"
	"github.com/bugg123/grpc-intro/serializer"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"

	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))
}

func TestJSONSerializer(t *testing.T) {
	binaryFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToJSONFile(laptop1, binaryFile)
	require.NoError(t, err)

}
