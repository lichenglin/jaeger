// Autogenerated by Thrift Compiler (0.13.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package agent

import (
	"bytes"
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/jaegertracing/jaeger/thrift-gen/jaeger"
	"github.com/jaegertracing/jaeger/thrift-gen/zipkincore"
	"reflect"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

var _ = jaeger.GoUnusedProtection__
var _ = zipkincore.GoUnusedProtection__

type Agent interface {
	// Parameters:
	//  - Spans
	EmitZipkinBatch(ctx context.Context, spans []*zipkincore.Span) (err error)
	// Parameters:
	//  - Batch
	EmitBatch(ctx context.Context, batch *jaeger.Batch) (err error)
}

type AgentClient struct {
	c thrift.TClient
}

func NewAgentClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *AgentClient {
	return &AgentClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewAgentClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *AgentClient {
	return &AgentClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewAgentClient(c thrift.TClient) *AgentClient {
	return &AgentClient{
		c: c,
	}
}

func (p *AgentClient) Client_() thrift.TClient {
	return p.c
}

// Parameters:
//  - Spans
func (p *AgentClient) EmitZipkinBatch(ctx context.Context, spans []*zipkincore.Span) (err error) {
	var _args0 AgentEmitZipkinBatchArgs
	_args0.Spans = spans
	if err := p.Client_().Call(ctx, "emitZipkinBatch", &_args0, nil); err != nil {
		return err
	}
	return nil
}

// Parameters:
//  - Batch
func (p *AgentClient) EmitBatch(ctx context.Context, batch *jaeger.Batch) (err error) {
	var _args1 AgentEmitBatchArgs
	_args1.Batch = batch
	if err := p.Client_().Call(ctx, "emitBatch", &_args1, nil); err != nil {
		return err
	}
	return nil
}

type AgentProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      Agent
}

func (p *AgentProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *AgentProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *AgentProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewAgentProcessor(handler Agent) *AgentProcessor {

	self2 := &AgentProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self2.processorMap["emitZipkinBatch"] = &agentProcessorEmitZipkinBatch{handler: handler}
	self2.processorMap["emitBatch"] = &agentProcessorEmitBatch{handler: handler}
	return self2
}

func (p *AgentProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x3 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x3.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x3

}

type agentProcessorEmitZipkinBatch struct {
	handler Agent
}

func (p *agentProcessorEmitZipkinBatch) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := AgentEmitZipkinBatchArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	if err2 = p.handler.EmitZipkinBatch(ctx, args.Spans); err2 != nil {
		return true, err2
	}
	return true, nil
}

type agentProcessorEmitBatch struct {
	handler Agent
}

func (p *agentProcessorEmitBatch) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := AgentEmitBatchArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	if err2 = p.handler.EmitBatch(ctx, args.Batch); err2 != nil {
		return true, err2
	}
	return true, nil
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Spans
type AgentEmitZipkinBatchArgs struct {
	Spans []*zipkincore.Span `thrift:"spans,1" db:"spans" json:"spans"`
}

func NewAgentEmitZipkinBatchArgs() *AgentEmitZipkinBatchArgs {
	return &AgentEmitZipkinBatchArgs{}
}

func (p *AgentEmitZipkinBatchArgs) GetSpans() []*zipkincore.Span {
	return p.Spans
}
func (p *AgentEmitZipkinBatchArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.LIST {
				if err := p.ReadField1(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AgentEmitZipkinBatchArgs) ReadField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*zipkincore.Span, 0, size)
	p.Spans = tSlice
	for i := 0; i < size; i++ {
		_elem4 := &zipkincore.Span{}
		if err := _elem4.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem4), err)
		}
		p.Spans = append(p.Spans, _elem4)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *AgentEmitZipkinBatchArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("emitZipkinBatch_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AgentEmitZipkinBatchArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("spans", thrift.LIST, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:spans: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Spans)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Spans {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:spans: ", p), err)
	}
	return err
}

func (p *AgentEmitZipkinBatchArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AgentEmitZipkinBatchArgs(%+v)", *p)
}

// Attributes:
//  - Batch
type AgentEmitBatchArgs struct {
	Batch *jaeger.Batch `thrift:"batch,1" db:"batch" json:"batch"`
}

func NewAgentEmitBatchArgs() *AgentEmitBatchArgs {
	return &AgentEmitBatchArgs{}
}

var AgentEmitBatchArgs_Batch_DEFAULT *jaeger.Batch

func (p *AgentEmitBatchArgs) GetBatch() *jaeger.Batch {
	if !p.IsSetBatch() {
		return AgentEmitBatchArgs_Batch_DEFAULT
	}
	return p.Batch
}
func (p *AgentEmitBatchArgs) IsSetBatch() bool {
	return p.Batch != nil
}

func (p *AgentEmitBatchArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AgentEmitBatchArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Batch = &jaeger.Batch{}
	if err := p.Batch.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Batch), err)
	}
	return nil
}

func (p *AgentEmitBatchArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("emitBatch_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AgentEmitBatchArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("batch", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:batch: ", p), err)
	}
	if err := p.Batch.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Batch), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:batch: ", p), err)
	}
	return err
}

func (p *AgentEmitBatchArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AgentEmitBatchArgs(%+v)", *p)
}
