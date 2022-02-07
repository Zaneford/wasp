package errors

import (
	"github.com/iotaledger/hive.go/marshalutil"
	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/kv/collections"
	"golang.org/x/xerrors"
)

// StateErrorCollectionWriter implements IErrorCollection. Is used for contract internal errors.
// It requires a reference to a KVStore such as the vmctx and the hname of the caller.
type StateErrorCollectionWriter struct {
	partition kv.KVStore
	hname     iscp.Hname
}

func NewStateErrorCollectionWriter(partition kv.KVStore, hname iscp.Hname) IErrorCollection {
	errorCollection := StateErrorCollectionWriter{
		partition: partition,
		hname:     hname,
	}

	return &errorCollection
}

func (e *StateErrorCollectionWriter) getErrorDefinitionMap() *collections.Map {
	mapName := ParamErrorDefinitionMap + string(e.hname.Bytes())

	return collections.NewMap(e.partition, mapName)
}

func (e *StateErrorCollectionWriter) Get(errorId uint16) (*ErrorDefinition, error) {
	errorMap := e.getErrorDefinitionMap()
	errorIdKey := codec.EncodeUint16(errorId)

	errorBytes, err := errorMap.GetAt(errorIdKey)

	if err != nil {
		return nil, err
	}

	errorDefinition, err := ErrorDefinitionFromMarshalUtil(marshalutil.New(errorBytes))

	if err != nil {
		return nil, err
	}

	return errorDefinition, nil
}

func (e *StateErrorCollectionWriter) Register(errorId uint16, messageFormat string) (*ErrorDefinition, error) {
	errorMap := e.getErrorDefinitionMap()
	mapKey := codec.EncodeUint16(errorId)

	if errorBytes, err := errorMap.GetAt(mapKey); err != nil {
		return nil, err
	} else if len(errorBytes) > 0 {
		return nil, xerrors.Errorf("Error already registered")
	}

	newError := ErrorDefinition{prefixId: e.hname, id: errorId, messageFormat: messageFormat}

	if err := errorMap.SetAt(mapKey, newError.Bytes()); err != nil {
		return nil, err
	}

	return &newError, nil
}

// StateErrorCollectionReader implements IErrorCollection partially. Is used for contract internal error readings only.
// It requires a reference to a KVStoreReader such as the vmctx and the hname of the caller.
type StateErrorCollectionReader struct {
	partition kv.KVStoreReader
	hname     iscp.Hname
}

func (e *StateErrorCollectionReader) getErrorDefinitionMap() *collections.ImmutableMap {
	mapName := ParamErrorDefinitionMap + string(e.hname.Bytes())

	return collections.NewMapReadOnly(e.partition, mapName)
}

func NewStateErrorCollectionReader(partition kv.KVStoreReader, hname iscp.Hname) IErrorCollection {
	errorCollection := StateErrorCollectionReader{
		partition: partition,
		hname:     hname,
	}

	return &errorCollection
}

func (e *StateErrorCollectionReader) Get(errorId uint16) (*ErrorDefinition, error) {
	errorMap := e.getErrorDefinitionMap()
	errorIdKey := codec.EncodeUint16(errorId)

	errorBytes, err := errorMap.GetAt(errorIdKey)

	if err != nil {
		return nil, err
	}

	errorDefinition, err := ErrorDefinitionFromMarshalUtil(marshalutil.New(errorBytes))

	if err != nil {
		return nil, err
	}

	return errorDefinition, nil
}

func (e *StateErrorCollectionReader) Register(errorId uint16, messageFormat string) (*ErrorDefinition, error) {
	return nil, xerrors.Errorf("Registering in read only maps is unsupported")
}
