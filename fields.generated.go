package field

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// AccountField is a STRING field.
type AccountField struct{ quickfix.FIXString }

// Tag returns tag.Account (1).
func (f AccountField) Tag() quickfix.Tag { return tag.Account }

// NewAccount returns a new AccountField initialized with val.
func NewAccount(val string) AccountField {
	return AccountField{quickfix.FIXString(val)}
}

func (f AccountField) Value() string { return f.String() }

// AggregatedBookField is a BOOLEAN field.
type AggregatedBookField struct{ quickfix.FIXBoolean }

// Tag returns tag.AggregatedBook (266).
func (f AggregatedBookField) Tag() quickfix.Tag { return tag.AggregatedBook }

// NewAggregatedBook returns a new AggregatedBookField initialized with val.
func NewAggregatedBook(val bool) AggregatedBookField {
	return AggregatedBookField{quickfix.FIXBoolean(val)}
}

func (f AggregatedBookField) Value() bool { return f.Bool() }

// ApplVerIDField is a enum.ApplVerID field.
type ApplVerIDField struct{ quickfix.FIXString }

// Tag returns tag.ApplVerID (1128).
func (f ApplVerIDField) Tag() quickfix.Tag { return tag.ApplVerID }

func NewApplVerID(val enum.ApplVerID) ApplVerIDField {
	return ApplVerIDField{quickfix.FIXString(val)}
}

func (f ApplVerIDField) Value() enum.ApplVerID { return enum.ApplVerID(f.String()) }

// AvgPxField is a PRICE field.
type AvgPxField struct{ quickfix.FIXDecimal }

// Tag returns tag.AvgPx (6).
func (f AvgPxField) Tag() quickfix.Tag { return tag.AvgPx }

// NewAvgPx returns a new AvgPxField initialized with val and scale.
func NewAvgPx(val decimal.Decimal, scale int32) AvgPxField {
	return AvgPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f AvgPxField) Value() (val decimal.Decimal) { return f.Decimal }

// BeginSeqNoField is a SEQNUM field.
type BeginSeqNoField struct{ quickfix.FIXInt }

// Tag returns tag.BeginSeqNo (7).
func (f BeginSeqNoField) Tag() quickfix.Tag { return tag.BeginSeqNo }

// NewBeginSeqNo returns a new BeginSeqNoField initialized with val.
func NewBeginSeqNo(val int) BeginSeqNoField {
	return BeginSeqNoField{quickfix.FIXInt(val)}
}

func (f BeginSeqNoField) Value() int { return f.Int() }

// BeginStringField is a STRING field.
type BeginStringField struct{ quickfix.FIXString }

// Tag returns tag.BeginString (8).
func (f BeginStringField) Tag() quickfix.Tag { return tag.BeginString }

// NewBeginString returns a new BeginStringField initialized with val.
func NewBeginString(val string) BeginStringField {
	return BeginStringField{quickfix.FIXString(val)}
}

func (f BeginStringField) Value() string { return f.String() }

// BidMDEntryIDField is a STRING field.
type BidMDEntryIDField struct{ quickfix.FIXString }

// Tag returns tag.BidMDEntryID (1745).
func (f BidMDEntryIDField) Tag() quickfix.Tag { return tag.BidMDEntryID }

// NewBidMDEntryID returns a new BidMDEntryIDField initialized with val.
func NewBidMDEntryID(val string) BidMDEntryIDField {
	return BidMDEntryIDField{quickfix.FIXString(val)}
}

func (f BidMDEntryIDField) Value() string { return f.String() }

// BidPxField is a PRICE field.
type BidPxField struct{ quickfix.FIXDecimal }

// Tag returns tag.BidPx (132).
func (f BidPxField) Tag() quickfix.Tag { return tag.BidPx }

// NewBidPx returns a new BidPxField initialized with val and scale.
func NewBidPx(val decimal.Decimal, scale int32) BidPxField {
	return BidPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f BidPxField) Value() (val decimal.Decimal) { return f.Decimal }

// BidQuoteIDField is a STRING field.
type BidQuoteIDField struct{ quickfix.FIXString }

// Tag returns tag.BidQuoteID (1747).
func (f BidQuoteIDField) Tag() quickfix.Tag { return tag.BidQuoteID }

// NewBidQuoteID returns a new BidQuoteIDField initialized with val.
func NewBidQuoteID(val string) BidQuoteIDField {
	return BidQuoteIDField{quickfix.FIXString(val)}
}

func (f BidQuoteIDField) Value() string { return f.String() }

// BidSizeField is a QTY field.
type BidSizeField struct{ quickfix.FIXDecimal }

// Tag returns tag.BidSize (134).
func (f BidSizeField) Tag() quickfix.Tag { return tag.BidSize }

// NewBidSize returns a new BidSizeField initialized with val and scale.
func NewBidSize(val decimal.Decimal, scale int32) BidSizeField {
	return BidSizeField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f BidSizeField) Value() (val decimal.Decimal) { return f.Decimal }

// BodyLengthField is a LENGTH field.
type BodyLengthField struct{ quickfix.FIXInt }

// Tag returns tag.BodyLength (9).
func (f BodyLengthField) Tag() quickfix.Tag { return tag.BodyLength }

// NewBodyLength returns a new BodyLengthField initialized with val.
func NewBodyLength(val int) BodyLengthField {
	return BodyLengthField{quickfix.FIXInt(val)}
}

func (f BodyLengthField) Value() int { return f.Int() }

// BusinessRejectReasonField is a enum.BusinessRejectReason field.
type BusinessRejectReasonField struct{ quickfix.FIXString }

// Tag returns tag.BusinessRejectReason (380).
func (f BusinessRejectReasonField) Tag() quickfix.Tag { return tag.BusinessRejectReason }

func NewBusinessRejectReason(val enum.BusinessRejectReason) BusinessRejectReasonField {
	return BusinessRejectReasonField{quickfix.FIXString(val)}
}

func (f BusinessRejectReasonField) Value() enum.BusinessRejectReason {
	return enum.BusinessRejectReason(f.String())
}

// CheckSumField is a STRING field.
type CheckSumField struct{ quickfix.FIXString }

// Tag returns tag.CheckSum (10).
func (f CheckSumField) Tag() quickfix.Tag { return tag.CheckSum }

// NewCheckSum returns a new CheckSumField initialized with val.
func NewCheckSum(val string) CheckSumField {
	return CheckSumField{quickfix.FIXString(val)}
}

func (f CheckSumField) Value() string { return f.String() }

// ClOrdIDField is a STRING field.
type ClOrdIDField struct{ quickfix.FIXString }

// Tag returns tag.ClOrdID (11).
func (f ClOrdIDField) Tag() quickfix.Tag { return tag.ClOrdID }

// NewClOrdID returns a new ClOrdIDField initialized with val.
func NewClOrdID(val string) ClOrdIDField {
	return ClOrdIDField{quickfix.FIXString(val)}
}

func (f ClOrdIDField) Value() string { return f.String() }

// CstmApplVerIDField is a STRING field.
type CstmApplVerIDField struct{ quickfix.FIXString }

// Tag returns tag.CstmApplVerID (1129).
func (f CstmApplVerIDField) Tag() quickfix.Tag { return tag.CstmApplVerID }

// NewCstmApplVerID returns a new CstmApplVerIDField initialized with val.
func NewCstmApplVerID(val string) CstmApplVerIDField {
	return CstmApplVerIDField{quickfix.FIXString(val)}
}

func (f CstmApplVerIDField) Value() string { return f.String() }

// CumQtyField is a QTY field.
type CumQtyField struct{ quickfix.FIXDecimal }

// Tag returns tag.CumQty (14).
func (f CumQtyField) Tag() quickfix.Tag { return tag.CumQty }

// NewCumQty returns a new CumQtyField initialized with val and scale.
func NewCumQty(val decimal.Decimal, scale int32) CumQtyField {
	return CumQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f CumQtyField) Value() (val decimal.Decimal) { return f.Decimal }

// CurrencyField is a CURRENCY field.
type CurrencyField struct{ quickfix.FIXString }

// Tag returns tag.Currency (15).
func (f CurrencyField) Tag() quickfix.Tag { return tag.Currency }

// NewCurrency returns a new CurrencyField initialized with val.
func NewCurrency(val string) CurrencyField {
	return CurrencyField{quickfix.FIXString(val)}
}

func (f CurrencyField) Value() string { return f.String() }

// CxlRejReasonField is a enum.CxlRejReason field.
type CxlRejReasonField struct{ quickfix.FIXString }

// Tag returns tag.CxlRejReason (102).
func (f CxlRejReasonField) Tag() quickfix.Tag { return tag.CxlRejReason }

func NewCxlRejReason(val enum.CxlRejReason) CxlRejReasonField {
	return CxlRejReasonField{quickfix.FIXString(val)}
}

func (f CxlRejReasonField) Value() enum.CxlRejReason { return enum.CxlRejReason(f.String()) }

// CxlRejResponseToField is a enum.CxlRejResponseTo field.
type CxlRejResponseToField struct{ quickfix.FIXString }

// Tag returns tag.CxlRejResponseTo (434).
func (f CxlRejResponseToField) Tag() quickfix.Tag { return tag.CxlRejResponseTo }

func NewCxlRejResponseTo(val enum.CxlRejResponseTo) CxlRejResponseToField {
	return CxlRejResponseToField{quickfix.FIXString(val)}
}

func (f CxlRejResponseToField) Value() enum.CxlRejResponseTo {
	return enum.CxlRejResponseTo(f.String())
}

// DefaultApplVerIDField is a STRING field.
type DefaultApplVerIDField struct{ quickfix.FIXString }

// Tag returns tag.DefaultApplVerID (1137).
func (f DefaultApplVerIDField) Tag() quickfix.Tag { return tag.DefaultApplVerID }

// NewDefaultApplVerID returns a new DefaultApplVerIDField initialized with val.
func NewDefaultApplVerID(val string) DefaultApplVerIDField {
	return DefaultApplVerIDField{quickfix.FIXString(val)}
}

func (f DefaultApplVerIDField) Value() string { return f.String() }

// DeliverToCompIDField is a STRING field.
type DeliverToCompIDField struct{ quickfix.FIXString }

// Tag returns tag.DeliverToCompID (128).
func (f DeliverToCompIDField) Tag() quickfix.Tag { return tag.DeliverToCompID }

// NewDeliverToCompID returns a new DeliverToCompIDField initialized with val.
func NewDeliverToCompID(val string) DeliverToCompIDField {
	return DeliverToCompIDField{quickfix.FIXString(val)}
}

func (f DeliverToCompIDField) Value() string { return f.String() }

// DeliverToLocationIDField is a STRING field.
type DeliverToLocationIDField struct{ quickfix.FIXString }

// Tag returns tag.DeliverToLocationID (145).
func (f DeliverToLocationIDField) Tag() quickfix.Tag { return tag.DeliverToLocationID }

// NewDeliverToLocationID returns a new DeliverToLocationIDField initialized with val.
func NewDeliverToLocationID(val string) DeliverToLocationIDField {
	return DeliverToLocationIDField{quickfix.FIXString(val)}
}

func (f DeliverToLocationIDField) Value() string { return f.String() }

// DeliverToSubIDField is a STRING field.
type DeliverToSubIDField struct{ quickfix.FIXString }

// Tag returns tag.DeliverToSubID (129).
func (f DeliverToSubIDField) Tag() quickfix.Tag { return tag.DeliverToSubID }

// NewDeliverToSubID returns a new DeliverToSubIDField initialized with val.
func NewDeliverToSubID(val string) DeliverToSubIDField {
	return DeliverToSubIDField{quickfix.FIXString(val)}
}

func (f DeliverToSubIDField) Value() string { return f.String() }

// EncodedTextField is a DATA field.
type EncodedTextField struct{ quickfix.FIXString }

// Tag returns tag.EncodedText (355).
func (f EncodedTextField) Tag() quickfix.Tag { return tag.EncodedText }

// NewEncodedText returns a new EncodedTextField initialized with val.
func NewEncodedText(val string) EncodedTextField {
	return EncodedTextField{quickfix.FIXString(val)}
}

func (f EncodedTextField) Value() string { return f.String() }

// EncodedTextLenField is a LENGTH field.
type EncodedTextLenField struct{ quickfix.FIXInt }

// Tag returns tag.EncodedTextLen (354).
func (f EncodedTextLenField) Tag() quickfix.Tag { return tag.EncodedTextLen }

// NewEncodedTextLen returns a new EncodedTextLenField initialized with val.
func NewEncodedTextLen(val int) EncodedTextLenField {
	return EncodedTextLenField{quickfix.FIXInt(val)}
}

func (f EncodedTextLenField) Value() int { return f.Int() }

// EncryptMethodField is a enum.EncryptMethod field.
type EncryptMethodField struct{ quickfix.FIXString }

// Tag returns tag.EncryptMethod (98).
func (f EncryptMethodField) Tag() quickfix.Tag { return tag.EncryptMethod }

func NewEncryptMethod(val enum.EncryptMethod) EncryptMethodField {
	return EncryptMethodField{quickfix.FIXString(val)}
}

func (f EncryptMethodField) Value() enum.EncryptMethod { return enum.EncryptMethod(f.String()) }

// EndSeqNoField is a SEQNUM field.
type EndSeqNoField struct{ quickfix.FIXInt }

// Tag returns tag.EndSeqNo (16).
func (f EndSeqNoField) Tag() quickfix.Tag { return tag.EndSeqNo }

// NewEndSeqNo returns a new EndSeqNoField initialized with val.
func NewEndSeqNo(val int) EndSeqNoField {
	return EndSeqNoField{quickfix.FIXInt(val)}
}

func (f EndSeqNoField) Value() int { return f.Int() }

// ExecIDField is a STRING field.
type ExecIDField struct{ quickfix.FIXString }

// Tag returns tag.ExecID (17).
func (f ExecIDField) Tag() quickfix.Tag { return tag.ExecID }

// NewExecID returns a new ExecIDField initialized with val.
func NewExecID(val string) ExecIDField {
	return ExecIDField{quickfix.FIXString(val)}
}

func (f ExecIDField) Value() string { return f.String() }

// ExecTypeField is a enum.ExecType field.
type ExecTypeField struct{ quickfix.FIXString }

// Tag returns tag.ExecType (150).
func (f ExecTypeField) Tag() quickfix.Tag { return tag.ExecType }

func NewExecType(val enum.ExecType) ExecTypeField {
	return ExecTypeField{quickfix.FIXString(val)}
}

func (f ExecTypeField) Value() enum.ExecType { return enum.ExecType(f.String()) }

// GapFillFlagField is a BOOLEAN field.
type GapFillFlagField struct{ quickfix.FIXBoolean }

// Tag returns tag.GapFillFlag (123).
func (f GapFillFlagField) Tag() quickfix.Tag { return tag.GapFillFlag }

// NewGapFillFlag returns a new GapFillFlagField initialized with val.
func NewGapFillFlag(val bool) GapFillFlagField {
	return GapFillFlagField{quickfix.FIXBoolean(val)}
}

func (f GapFillFlagField) Value() bool { return f.Bool() }

// HeadlineField is a STRING field.
type HeadlineField struct{ quickfix.FIXString }

// Tag returns tag.Headline (148).
func (f HeadlineField) Tag() quickfix.Tag { return tag.Headline }

// NewHeadline returns a new HeadlineField initialized with val.
func NewHeadline(val string) HeadlineField {
	return HeadlineField{quickfix.FIXString(val)}
}

func (f HeadlineField) Value() string { return f.String() }

// HeartBtIntField is a INT field.
type HeartBtIntField struct{ quickfix.FIXInt }

// Tag returns tag.HeartBtInt (108).
func (f HeartBtIntField) Tag() quickfix.Tag { return tag.HeartBtInt }

// NewHeartBtInt returns a new HeartBtIntField initialized with val.
func NewHeartBtInt(val int) HeartBtIntField {
	return HeartBtIntField{quickfix.FIXInt(val)}
}

func (f HeartBtIntField) Value() int { return f.Int() }

// HopCompIDField is a STRING field.
type HopCompIDField struct{ quickfix.FIXString }

// Tag returns tag.HopCompID (628).
func (f HopCompIDField) Tag() quickfix.Tag { return tag.HopCompID }

// NewHopCompID returns a new HopCompIDField initialized with val.
func NewHopCompID(val string) HopCompIDField {
	return HopCompIDField{quickfix.FIXString(val)}
}

func (f HopCompIDField) Value() string { return f.String() }

// HopRefIDField is a SEQNUM field.
type HopRefIDField struct{ quickfix.FIXInt }

// Tag returns tag.HopRefID (630).
func (f HopRefIDField) Tag() quickfix.Tag { return tag.HopRefID }

// NewHopRefID returns a new HopRefIDField initialized with val.
func NewHopRefID(val int) HopRefIDField {
	return HopRefIDField{quickfix.FIXInt(val)}
}

func (f HopRefIDField) Value() int { return f.Int() }

// HopSendingTimeField is a UTCTIMESTAMP field.
type HopSendingTimeField struct{ quickfix.FIXUTCTimestamp }

// Tag returns tag.HopSendingTime (629).
func (f HopSendingTimeField) Tag() quickfix.Tag { return tag.HopSendingTime }

// NewHopSendingTime returns a new HopSendingTimeField initialized with val.
func NewHopSendingTime(val time.Time) HopSendingTimeField {
	return NewHopSendingTimeWithPrecision(val, quickfix.Millis)
}

// NewHopSendingTimeNoMillis returns a new HopSendingTimeField initialized with val without millisecs.
func NewHopSendingTimeNoMillis(val time.Time) HopSendingTimeField {
	return NewHopSendingTimeWithPrecision(val, quickfix.Seconds)
}

// NewHopSendingTimeWithPrecision returns a new HopSendingTimeField initialized with val of specified precision.
func NewHopSendingTimeWithPrecision(val time.Time, precision quickfix.TimestampPrecision) HopSendingTimeField {
	return HopSendingTimeField{quickfix.FIXUTCTimestamp{Time: val, Precision: precision}}
}

func (f HopSendingTimeField) Value() time.Time { return f.Time }

// LastFragmentField is a BOOLEAN field.
type LastFragmentField struct{ quickfix.FIXBoolean }

// Tag returns tag.LastFragment (893).
func (f LastFragmentField) Tag() quickfix.Tag { return tag.LastFragment }

// NewLastFragment returns a new LastFragmentField initialized with val.
func NewLastFragment(val bool) LastFragmentField {
	return LastFragmentField{quickfix.FIXBoolean(val)}
}

func (f LastFragmentField) Value() bool { return f.Bool() }

// LastMsgSeqNumProcessedField is a SEQNUM field.
type LastMsgSeqNumProcessedField struct{ quickfix.FIXInt }

// Tag returns tag.LastMsgSeqNumProcessed (369).
func (f LastMsgSeqNumProcessedField) Tag() quickfix.Tag { return tag.LastMsgSeqNumProcessed }

// NewLastMsgSeqNumProcessed returns a new LastMsgSeqNumProcessedField initialized with val.
func NewLastMsgSeqNumProcessed(val int) LastMsgSeqNumProcessedField {
	return LastMsgSeqNumProcessedField{quickfix.FIXInt(val)}
}

func (f LastMsgSeqNumProcessedField) Value() int { return f.Int() }

// LastPxField is a PRICE field.
type LastPxField struct{ quickfix.FIXDecimal }

// Tag returns tag.LastPx (31).
func (f LastPxField) Tag() quickfix.Tag { return tag.LastPx }

// NewLastPx returns a new LastPxField initialized with val and scale.
func NewLastPx(val decimal.Decimal, scale int32) LastPxField {
	return LastPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LastPxField) Value() (val decimal.Decimal) { return f.Decimal }

// LastQtyField is a QTY field.
type LastQtyField struct{ quickfix.FIXDecimal }

// Tag returns tag.LastQty (32).
func (f LastQtyField) Tag() quickfix.Tag { return tag.LastQty }

// NewLastQty returns a new LastQtyField initialized with val and scale.
func NewLastQty(val decimal.Decimal, scale int32) LastQtyField {
	return LastQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LastQtyField) Value() (val decimal.Decimal) { return f.Decimal }

// LastRptRequestedField is a BOOLEAN field.
type LastRptRequestedField struct{ quickfix.FIXBoolean }

// Tag returns tag.LastRptRequested (912).
func (f LastRptRequestedField) Tag() quickfix.Tag { return tag.LastRptRequested }

// NewLastRptRequested returns a new LastRptRequestedField initialized with val.
func NewLastRptRequested(val bool) LastRptRequestedField {
	return LastRptRequestedField{quickfix.FIXBoolean(val)}
}

func (f LastRptRequestedField) Value() bool { return f.Bool() }

// LastUpdateTimeField is a UTCTIMESTAMP field.
type LastUpdateTimeField struct{ quickfix.FIXUTCTimestamp }

// Tag returns tag.LastUpdateTime (779).
func (f LastUpdateTimeField) Tag() quickfix.Tag { return tag.LastUpdateTime }

// NewLastUpdateTime returns a new LastUpdateTimeField initialized with val.
func NewLastUpdateTime(val time.Time) LastUpdateTimeField {
	return NewLastUpdateTimeWithPrecision(val, quickfix.Millis)
}

// NewLastUpdateTimeNoMillis returns a new LastUpdateTimeField initialized with val without millisecs.
func NewLastUpdateTimeNoMillis(val time.Time) LastUpdateTimeField {
	return NewLastUpdateTimeWithPrecision(val, quickfix.Seconds)
}

// NewLastUpdateTimeWithPrecision returns a new LastUpdateTimeField initialized with val of specified precision.
func NewLastUpdateTimeWithPrecision(val time.Time, precision quickfix.TimestampPrecision) LastUpdateTimeField {
	return LastUpdateTimeField{quickfix.FIXUTCTimestamp{Time: val, Precision: precision}}
}

func (f LastUpdateTimeField) Value() time.Time { return f.Time }

// LeavesQtyField is a QTY field.
type LeavesQtyField struct{ quickfix.FIXDecimal }

// Tag returns tag.LeavesQty (151).
func (f LeavesQtyField) Tag() quickfix.Tag { return tag.LeavesQty }

// NewLeavesQty returns a new LeavesQtyField initialized with val and scale.
func NewLeavesQty(val decimal.Decimal, scale int32) LeavesQtyField {
	return LeavesQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f LeavesQtyField) Value() (val decimal.Decimal) { return f.Decimal }

// MDEntryDateField is a UTCDATEONLY field.
type MDEntryDateField struct{ quickfix.FIXString }

// Tag returns tag.MDEntryDate (272).
func (f MDEntryDateField) Tag() quickfix.Tag { return tag.MDEntryDate }

// NewMDEntryDate returns a new MDEntryDateField initialized with val.
func NewMDEntryDate(val string) MDEntryDateField {
	return MDEntryDateField{quickfix.FIXString(val)}
}

func (f MDEntryDateField) Value() string { return f.String() }

// MDEntryIDField is a STRING field.
type MDEntryIDField struct{ quickfix.FIXString }

// Tag returns tag.MDEntryID (278).
func (f MDEntryIDField) Tag() quickfix.Tag { return tag.MDEntryID }

// NewMDEntryID returns a new MDEntryIDField initialized with val.
func NewMDEntryID(val string) MDEntryIDField {
	return MDEntryIDField{quickfix.FIXString(val)}
}

func (f MDEntryIDField) Value() string { return f.String() }

// MDEntryPxField is a PRICE field.
type MDEntryPxField struct{ quickfix.FIXDecimal }

// Tag returns tag.MDEntryPx (270).
func (f MDEntryPxField) Tag() quickfix.Tag { return tag.MDEntryPx }

// NewMDEntryPx returns a new MDEntryPxField initialized with val and scale.
func NewMDEntryPx(val decimal.Decimal, scale int32) MDEntryPxField {
	return MDEntryPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MDEntryPxField) Value() (val decimal.Decimal) { return f.Decimal }

// MDEntryRefIDField is a STRING field.
type MDEntryRefIDField struct{ quickfix.FIXString }

// Tag returns tag.MDEntryRefID (280).
func (f MDEntryRefIDField) Tag() quickfix.Tag { return tag.MDEntryRefID }

// NewMDEntryRefID returns a new MDEntryRefIDField initialized with val.
func NewMDEntryRefID(val string) MDEntryRefIDField {
	return MDEntryRefIDField{quickfix.FIXString(val)}
}

func (f MDEntryRefIDField) Value() string { return f.String() }

// MDEntrySizeField is a QTY field.
type MDEntrySizeField struct{ quickfix.FIXDecimal }

// Tag returns tag.MDEntrySize (271).
func (f MDEntrySizeField) Tag() quickfix.Tag { return tag.MDEntrySize }

// NewMDEntrySize returns a new MDEntrySizeField initialized with val and scale.
func NewMDEntrySize(val decimal.Decimal, scale int32) MDEntrySizeField {
	return MDEntrySizeField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f MDEntrySizeField) Value() (val decimal.Decimal) { return f.Decimal }

// MDEntryTimeField is a UTCTIMEONLY field.
type MDEntryTimeField struct{ quickfix.FIXString }

// Tag returns tag.MDEntryTime (273).
func (f MDEntryTimeField) Tag() quickfix.Tag { return tag.MDEntryTime }

// NewMDEntryTime returns a new MDEntryTimeField initialized with val.
func NewMDEntryTime(val string) MDEntryTimeField {
	return MDEntryTimeField{quickfix.FIXString(val)}
}

func (f MDEntryTimeField) Value() string { return f.String() }

// MDEntryTypeField is a enum.MDEntryType field.
type MDEntryTypeField struct{ quickfix.FIXString }

// Tag returns tag.MDEntryType (269).
func (f MDEntryTypeField) Tag() quickfix.Tag { return tag.MDEntryType }

func NewMDEntryType(val enum.MDEntryType) MDEntryTypeField {
	return MDEntryTypeField{quickfix.FIXString(val)}
}

func (f MDEntryTypeField) Value() enum.MDEntryType { return enum.MDEntryType(f.String()) }

// MDReqIDField is a STRING field.
type MDReqIDField struct{ quickfix.FIXString }

// Tag returns tag.MDReqID (262).
func (f MDReqIDField) Tag() quickfix.Tag { return tag.MDReqID }

// NewMDReqID returns a new MDReqIDField initialized with val.
func NewMDReqID(val string) MDReqIDField {
	return MDReqIDField{quickfix.FIXString(val)}
}

func (f MDReqIDField) Value() string { return f.String() }

// MDReqRejReasonField is a enum.MDReqRejReason field.
type MDReqRejReasonField struct{ quickfix.FIXString }

// Tag returns tag.MDReqRejReason (281).
func (f MDReqRejReasonField) Tag() quickfix.Tag { return tag.MDReqRejReason }

func NewMDReqRejReason(val enum.MDReqRejReason) MDReqRejReasonField {
	return MDReqRejReasonField{quickfix.FIXString(val)}
}

func (f MDReqRejReasonField) Value() enum.MDReqRejReason { return enum.MDReqRejReason(f.String()) }

// MDUpdateActionField is a enum.MDUpdateAction field.
type MDUpdateActionField struct{ quickfix.FIXString }

// Tag returns tag.MDUpdateAction (279).
func (f MDUpdateActionField) Tag() quickfix.Tag { return tag.MDUpdateAction }

func NewMDUpdateAction(val enum.MDUpdateAction) MDUpdateActionField {
	return MDUpdateActionField{quickfix.FIXString(val)}
}

func (f MDUpdateActionField) Value() enum.MDUpdateAction { return enum.MDUpdateAction(f.String()) }

// MDUpdateTypeField is a enum.MDUpdateType field.
type MDUpdateTypeField struct{ quickfix.FIXString }

// Tag returns tag.MDUpdateType (265).
func (f MDUpdateTypeField) Tag() quickfix.Tag { return tag.MDUpdateType }

func NewMDUpdateType(val enum.MDUpdateType) MDUpdateTypeField {
	return MDUpdateTypeField{quickfix.FIXString(val)}
}

func (f MDUpdateTypeField) Value() enum.MDUpdateType { return enum.MDUpdateType(f.String()) }

// MarketDepthField is a INT field.
type MarketDepthField struct{ quickfix.FIXInt }

// Tag returns tag.MarketDepth (264).
func (f MarketDepthField) Tag() quickfix.Tag { return tag.MarketDepth }

// NewMarketDepth returns a new MarketDepthField initialized with val.
func NewMarketDepth(val int) MarketDepthField {
	return MarketDepthField{quickfix.FIXInt(val)}
}

func (f MarketDepthField) Value() int { return f.Int() }

// MarketIDField is a EXCHANGE field.
type MarketIDField struct{ quickfix.FIXString }

// Tag returns tag.MarketID (1301).
func (f MarketIDField) Tag() quickfix.Tag { return tag.MarketID }

// NewMarketID returns a new MarketIDField initialized with val.
func NewMarketID(val string) MarketIDField {
	return MarketIDField{quickfix.FIXString(val)}
}

func (f MarketIDField) Value() string { return f.String() }

// MassActionReportIDField is a STRING field.
type MassActionReportIDField struct{ quickfix.FIXString }

// Tag returns tag.MassActionReportID (1369).
func (f MassActionReportIDField) Tag() quickfix.Tag { return tag.MassActionReportID }

// NewMassActionReportID returns a new MassActionReportIDField initialized with val.
func NewMassActionReportID(val string) MassActionReportIDField {
	return MassActionReportIDField{quickfix.FIXString(val)}
}

func (f MassActionReportIDField) Value() string { return f.String() }

// MassCancelRejectReasonField is a enum.MassCancelRejectReason field.
type MassCancelRejectReasonField struct{ quickfix.FIXString }

// Tag returns tag.MassCancelRejectReason (532).
func (f MassCancelRejectReasonField) Tag() quickfix.Tag { return tag.MassCancelRejectReason }

func NewMassCancelRejectReason(val enum.MassCancelRejectReason) MassCancelRejectReasonField {
	return MassCancelRejectReasonField{quickfix.FIXString(val)}
}

func (f MassCancelRejectReasonField) Value() enum.MassCancelRejectReason {
	return enum.MassCancelRejectReason(f.String())
}

// MassCancelRequestTypeField is a enum.MassCancelRequestType field.
type MassCancelRequestTypeField struct{ quickfix.FIXString }

// Tag returns tag.MassCancelRequestType (530).
func (f MassCancelRequestTypeField) Tag() quickfix.Tag { return tag.MassCancelRequestType }

func NewMassCancelRequestType(val enum.MassCancelRequestType) MassCancelRequestTypeField {
	return MassCancelRequestTypeField{quickfix.FIXString(val)}
}

func (f MassCancelRequestTypeField) Value() enum.MassCancelRequestType {
	return enum.MassCancelRequestType(f.String())
}

// MassCancelResponseField is a enum.MassCancelResponse field.
type MassCancelResponseField struct{ quickfix.FIXString }

// Tag returns tag.MassCancelResponse (531).
func (f MassCancelResponseField) Tag() quickfix.Tag { return tag.MassCancelResponse }

func NewMassCancelResponse(val enum.MassCancelResponse) MassCancelResponseField {
	return MassCancelResponseField{quickfix.FIXString(val)}
}

func (f MassCancelResponseField) Value() enum.MassCancelResponse {
	return enum.MassCancelResponse(f.String())
}

// MassStatusReqIDField is a STRING field.
type MassStatusReqIDField struct{ quickfix.FIXString }

// Tag returns tag.MassStatusReqID (584).
func (f MassStatusReqIDField) Tag() quickfix.Tag { return tag.MassStatusReqID }

// NewMassStatusReqID returns a new MassStatusReqIDField initialized with val.
func NewMassStatusReqID(val string) MassStatusReqIDField {
	return MassStatusReqIDField{quickfix.FIXString(val)}
}

func (f MassStatusReqIDField) Value() string { return f.String() }

// MassStatusReqTypeField is a enum.MassStatusReqType field.
type MassStatusReqTypeField struct{ quickfix.FIXString }

// Tag returns tag.MassStatusReqType (585).
func (f MassStatusReqTypeField) Tag() quickfix.Tag { return tag.MassStatusReqType }

func NewMassStatusReqType(val enum.MassStatusReqType) MassStatusReqTypeField {
	return MassStatusReqTypeField{quickfix.FIXString(val)}
}

func (f MassStatusReqTypeField) Value() enum.MassStatusReqType {
	return enum.MassStatusReqType(f.String())
}

// MaxMessageSizeField is a LENGTH field.
type MaxMessageSizeField struct{ quickfix.FIXInt }

// Tag returns tag.MaxMessageSize (383).
func (f MaxMessageSizeField) Tag() quickfix.Tag { return tag.MaxMessageSize }

// NewMaxMessageSize returns a new MaxMessageSizeField initialized with val.
func NewMaxMessageSize(val int) MaxMessageSizeField {
	return MaxMessageSizeField{quickfix.FIXInt(val)}
}

func (f MaxMessageSizeField) Value() int { return f.Int() }

// MessageEncodingField is a enum.MessageEncoding field.
type MessageEncodingField struct{ quickfix.FIXString }

// Tag returns tag.MessageEncoding (347).
func (f MessageEncodingField) Tag() quickfix.Tag { return tag.MessageEncoding }

func NewMessageEncoding(val enum.MessageEncoding) MessageEncodingField {
	return MessageEncodingField{quickfix.FIXString(val)}
}

func (f MessageEncodingField) Value() enum.MessageEncoding { return enum.MessageEncoding(f.String()) }

// MsgDirectionField is a enum.MsgDirection field.
type MsgDirectionField struct{ quickfix.FIXString }

// Tag returns tag.MsgDirection (385).
func (f MsgDirectionField) Tag() quickfix.Tag { return tag.MsgDirection }

func NewMsgDirection(val enum.MsgDirection) MsgDirectionField {
	return MsgDirectionField{quickfix.FIXString(val)}
}

func (f MsgDirectionField) Value() enum.MsgDirection { return enum.MsgDirection(f.String()) }

// MsgSeqNumField is a SEQNUM field.
type MsgSeqNumField struct{ quickfix.FIXInt }

// Tag returns tag.MsgSeqNum (34).
func (f MsgSeqNumField) Tag() quickfix.Tag { return tag.MsgSeqNum }

// NewMsgSeqNum returns a new MsgSeqNumField initialized with val.
func NewMsgSeqNum(val int) MsgSeqNumField {
	return MsgSeqNumField{quickfix.FIXInt(val)}
}

func (f MsgSeqNumField) Value() int { return f.Int() }

// MsgTypeField is a enum.MsgType field.
type MsgTypeField struct{ quickfix.FIXString }

// Tag returns tag.MsgType (35).
func (f MsgTypeField) Tag() quickfix.Tag { return tag.MsgType }

func NewMsgType(val enum.MsgType) MsgTypeField {
	return MsgTypeField{quickfix.FIXString(val)}
}

func (f MsgTypeField) Value() enum.MsgType { return enum.MsgType(f.String()) }

// NewSeqNoField is a SEQNUM field.
type NewSeqNoField struct{ quickfix.FIXInt }

// Tag returns tag.NewSeqNo (36).
func (f NewSeqNoField) Tag() quickfix.Tag { return tag.NewSeqNo }

// NewNewSeqNo returns a new NewSeqNoField initialized with val.
func NewNewSeqNo(val int) NewSeqNoField {
	return NewSeqNoField{quickfix.FIXInt(val)}
}

func (f NewSeqNoField) Value() int { return f.Int() }

// NextExpectedMsgSeqNumField is a SEQNUM field.
type NextExpectedMsgSeqNumField struct{ quickfix.FIXInt }

// Tag returns tag.NextExpectedMsgSeqNum (789).
func (f NextExpectedMsgSeqNumField) Tag() quickfix.Tag { return tag.NextExpectedMsgSeqNum }

// NewNextExpectedMsgSeqNum returns a new NextExpectedMsgSeqNumField initialized with val.
func NewNextExpectedMsgSeqNum(val int) NextExpectedMsgSeqNumField {
	return NextExpectedMsgSeqNumField{quickfix.FIXInt(val)}
}

func (f NextExpectedMsgSeqNumField) Value() int { return f.Int() }

// NoHopsField is a NUMINGROUP field.
type NoHopsField struct{ quickfix.FIXInt }

// Tag returns tag.NoHops (627).
func (f NoHopsField) Tag() quickfix.Tag { return tag.NoHops }

// NewNoHops returns a new NoHopsField initialized with val.
func NewNoHops(val int) NoHopsField {
	return NoHopsField{quickfix.FIXInt(val)}
}

func (f NoHopsField) Value() int { return f.Int() }

// NoLinesOfTextField is a NUMINGROUP field.
type NoLinesOfTextField struct{ quickfix.FIXInt }

// Tag returns tag.NoLinesOfText (33).
func (f NoLinesOfTextField) Tag() quickfix.Tag { return tag.NoLinesOfText }

// NewNoLinesOfText returns a new NoLinesOfTextField initialized with val.
func NewNoLinesOfText(val int) NoLinesOfTextField {
	return NoLinesOfTextField{quickfix.FIXInt(val)}
}

func (f NoLinesOfTextField) Value() int { return f.Int() }

// NoMDEntriesField is a NUMINGROUP field.
type NoMDEntriesField struct{ quickfix.FIXInt }

// Tag returns tag.NoMDEntries (268).
func (f NoMDEntriesField) Tag() quickfix.Tag { return tag.NoMDEntries }

// NewNoMDEntries returns a new NoMDEntriesField initialized with val.
func NewNoMDEntries(val int) NoMDEntriesField {
	return NoMDEntriesField{quickfix.FIXInt(val)}
}

func (f NoMDEntriesField) Value() int { return f.Int() }

// NoMDEntryTypesField is a NUMINGROUP field.
type NoMDEntryTypesField struct{ quickfix.FIXInt }

// Tag returns tag.NoMDEntryTypes (267).
func (f NoMDEntryTypesField) Tag() quickfix.Tag { return tag.NoMDEntryTypes }

// NewNoMDEntryTypes returns a new NoMDEntryTypesField initialized with val.
func NewNoMDEntryTypes(val int) NoMDEntryTypesField {
	return NoMDEntryTypesField{quickfix.FIXInt(val)}
}

func (f NoMDEntryTypesField) Value() int { return f.Int() }

// NoMsgTypesField is a NUMINGROUP field.
type NoMsgTypesField struct{ quickfix.FIXInt }

// Tag returns tag.NoMsgTypes (384).
func (f NoMsgTypesField) Tag() quickfix.Tag { return tag.NoMsgTypes }

// NewNoMsgTypes returns a new NoMsgTypesField initialized with val.
func NewNoMsgTypes(val int) NoMsgTypesField {
	return NoMsgTypesField{quickfix.FIXInt(val)}
}

func (f NoMsgTypesField) Value() int { return f.Int() }

// NoPartyIDsField is a NUMINGROUP field.
type NoPartyIDsField struct{ quickfix.FIXInt }

// Tag returns tag.NoPartyIDs (453).
func (f NoPartyIDsField) Tag() quickfix.Tag { return tag.NoPartyIDs }

// NewNoPartyIDs returns a new NoPartyIDsField initialized with val.
func NewNoPartyIDs(val int) NoPartyIDsField {
	return NoPartyIDsField{quickfix.FIXInt(val)}
}

func (f NoPartyIDsField) Value() int { return f.Int() }

// NoPartySubIDsField is a NUMINGROUP field.
type NoPartySubIDsField struct{ quickfix.FIXInt }

// Tag returns tag.NoPartySubIDs (802).
func (f NoPartySubIDsField) Tag() quickfix.Tag { return tag.NoPartySubIDs }

// NewNoPartySubIDs returns a new NoPartySubIDsField initialized with val.
func NewNoPartySubIDs(val int) NoPartySubIDsField {
	return NoPartySubIDsField{quickfix.FIXInt(val)}
}

func (f NoPartySubIDsField) Value() int { return f.Int() }

// NoQuoteEntriesField is a NUMINGROUP field.
type NoQuoteEntriesField struct{ quickfix.FIXInt }

// Tag returns tag.NoQuoteEntries (295).
func (f NoQuoteEntriesField) Tag() quickfix.Tag { return tag.NoQuoteEntries }

// NewNoQuoteEntries returns a new NoQuoteEntriesField initialized with val.
func NewNoQuoteEntries(val int) NoQuoteEntriesField {
	return NoQuoteEntriesField{quickfix.FIXInt(val)}
}

func (f NoQuoteEntriesField) Value() int { return f.Int() }

// NoRelatedSymField is a NUMINGROUP field.
type NoRelatedSymField struct{ quickfix.FIXInt }

// Tag returns tag.NoRelatedSym (146).
func (f NoRelatedSymField) Tag() quickfix.Tag { return tag.NoRelatedSym }

// NewNoRelatedSym returns a new NoRelatedSymField initialized with val.
func NewNoRelatedSym(val int) NoRelatedSymField {
	return NoRelatedSymField{quickfix.FIXInt(val)}
}

func (f NoRelatedSymField) Value() int { return f.Int() }

// OfferMDEntryIDField is a STRING field.
type OfferMDEntryIDField struct{ quickfix.FIXString }

// Tag returns tag.OfferMDEntryID (1746).
func (f OfferMDEntryIDField) Tag() quickfix.Tag { return tag.OfferMDEntryID }

// NewOfferMDEntryID returns a new OfferMDEntryIDField initialized with val.
func NewOfferMDEntryID(val string) OfferMDEntryIDField {
	return OfferMDEntryIDField{quickfix.FIXString(val)}
}

func (f OfferMDEntryIDField) Value() string { return f.String() }

// OfferPxField is a PRICE field.
type OfferPxField struct{ quickfix.FIXDecimal }

// Tag returns tag.OfferPx (133).
func (f OfferPxField) Tag() quickfix.Tag { return tag.OfferPx }

// NewOfferPx returns a new OfferPxField initialized with val and scale.
func NewOfferPx(val decimal.Decimal, scale int32) OfferPxField {
	return OfferPxField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f OfferPxField) Value() (val decimal.Decimal) { return f.Decimal }

// OfferQuoteIDField is a STRING field.
type OfferQuoteIDField struct{ quickfix.FIXString }

// Tag returns tag.OfferQuoteID (1748).
func (f OfferQuoteIDField) Tag() quickfix.Tag { return tag.OfferQuoteID }

// NewOfferQuoteID returns a new OfferQuoteIDField initialized with val.
func NewOfferQuoteID(val string) OfferQuoteIDField {
	return OfferQuoteIDField{quickfix.FIXString(val)}
}

func (f OfferQuoteIDField) Value() string { return f.String() }

// OfferSizeField is a QTY field.
type OfferSizeField struct{ quickfix.FIXDecimal }

// Tag returns tag.OfferSize (135).
func (f OfferSizeField) Tag() quickfix.Tag { return tag.OfferSize }

// NewOfferSize returns a new OfferSizeField initialized with val and scale.
func NewOfferSize(val decimal.Decimal, scale int32) OfferSizeField {
	return OfferSizeField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f OfferSizeField) Value() (val decimal.Decimal) { return f.Decimal }

// OnBehalfOfCompIDField is a STRING field.
type OnBehalfOfCompIDField struct{ quickfix.FIXString }

// Tag returns tag.OnBehalfOfCompID (115).
func (f OnBehalfOfCompIDField) Tag() quickfix.Tag { return tag.OnBehalfOfCompID }

// NewOnBehalfOfCompID returns a new OnBehalfOfCompIDField initialized with val.
func NewOnBehalfOfCompID(val string) OnBehalfOfCompIDField {
	return OnBehalfOfCompIDField{quickfix.FIXString(val)}
}

func (f OnBehalfOfCompIDField) Value() string { return f.String() }

// OnBehalfOfLocationIDField is a STRING field.
type OnBehalfOfLocationIDField struct{ quickfix.FIXString }

// Tag returns tag.OnBehalfOfLocationID (144).
func (f OnBehalfOfLocationIDField) Tag() quickfix.Tag { return tag.OnBehalfOfLocationID }

// NewOnBehalfOfLocationID returns a new OnBehalfOfLocationIDField initialized with val.
func NewOnBehalfOfLocationID(val string) OnBehalfOfLocationIDField {
	return OnBehalfOfLocationIDField{quickfix.FIXString(val)}
}

func (f OnBehalfOfLocationIDField) Value() string { return f.String() }

// OnBehalfOfSubIDField is a STRING field.
type OnBehalfOfSubIDField struct{ quickfix.FIXString }

// Tag returns tag.OnBehalfOfSubID (116).
func (f OnBehalfOfSubIDField) Tag() quickfix.Tag { return tag.OnBehalfOfSubID }

// NewOnBehalfOfSubID returns a new OnBehalfOfSubIDField initialized with val.
func NewOnBehalfOfSubID(val string) OnBehalfOfSubIDField {
	return OnBehalfOfSubIDField{quickfix.FIXString(val)}
}

func (f OnBehalfOfSubIDField) Value() string { return f.String() }

// OpenCloseSettlFlagField is a enum.OpenCloseSettlFlag field.
type OpenCloseSettlFlagField struct{ quickfix.FIXString }

// Tag returns tag.OpenCloseSettlFlag (286).
func (f OpenCloseSettlFlagField) Tag() quickfix.Tag { return tag.OpenCloseSettlFlag }

func NewOpenCloseSettlFlag(val enum.OpenCloseSettlFlag) OpenCloseSettlFlagField {
	return OpenCloseSettlFlagField{quickfix.FIXString(val)}
}

func (f OpenCloseSettlFlagField) Value() enum.OpenCloseSettlFlag {
	return enum.OpenCloseSettlFlag(f.String())
}

// OrdRejReasonField is a enum.OrdRejReason field.
type OrdRejReasonField struct{ quickfix.FIXString }

// Tag returns tag.OrdRejReason (103).
func (f OrdRejReasonField) Tag() quickfix.Tag { return tag.OrdRejReason }

func NewOrdRejReason(val enum.OrdRejReason) OrdRejReasonField {
	return OrdRejReasonField{quickfix.FIXString(val)}
}

func (f OrdRejReasonField) Value() enum.OrdRejReason { return enum.OrdRejReason(f.String()) }

// OrdStatusField is a enum.OrdStatus field.
type OrdStatusField struct{ quickfix.FIXString }

// Tag returns tag.OrdStatus (39).
func (f OrdStatusField) Tag() quickfix.Tag { return tag.OrdStatus }

func NewOrdStatus(val enum.OrdStatus) OrdStatusField {
	return OrdStatusField{quickfix.FIXString(val)}
}

func (f OrdStatusField) Value() enum.OrdStatus { return enum.OrdStatus(f.String()) }

// OrdStatusReqIDField is a STRING field.
type OrdStatusReqIDField struct{ quickfix.FIXString }

// Tag returns tag.OrdStatusReqID (790).
func (f OrdStatusReqIDField) Tag() quickfix.Tag { return tag.OrdStatusReqID }

// NewOrdStatusReqID returns a new OrdStatusReqIDField initialized with val.
func NewOrdStatusReqID(val string) OrdStatusReqIDField {
	return OrdStatusReqIDField{quickfix.FIXString(val)}
}

func (f OrdStatusReqIDField) Value() string { return f.String() }

// OrdTypeField is a enum.OrdType field.
type OrdTypeField struct{ quickfix.FIXString }

// Tag returns tag.OrdType (40).
func (f OrdTypeField) Tag() quickfix.Tag { return tag.OrdType }

func NewOrdType(val enum.OrdType) OrdTypeField {
	return OrdTypeField{quickfix.FIXString(val)}
}

func (f OrdTypeField) Value() enum.OrdType { return enum.OrdType(f.String()) }

// OrderIDField is a STRING field.
type OrderIDField struct{ quickfix.FIXString }

// Tag returns tag.OrderID (37).
func (f OrderIDField) Tag() quickfix.Tag { return tag.OrderID }

// NewOrderID returns a new OrderIDField initialized with val.
func NewOrderID(val string) OrderIDField {
	return OrderIDField{quickfix.FIXString(val)}
}

func (f OrderIDField) Value() string { return f.String() }

// OrderOriginationField is a enum.OrderOrigination field.
type OrderOriginationField struct{ quickfix.FIXString }

// Tag returns tag.OrderOrigination (1724).
func (f OrderOriginationField) Tag() quickfix.Tag { return tag.OrderOrigination }

func NewOrderOrigination(val enum.OrderOrigination) OrderOriginationField {
	return OrderOriginationField{quickfix.FIXString(val)}
}

func (f OrderOriginationField) Value() enum.OrderOrigination {
	return enum.OrderOrigination(f.String())
}

// OrderQtyField is a QTY field.
type OrderQtyField struct{ quickfix.FIXDecimal }

// Tag returns tag.OrderQty (38).
func (f OrderQtyField) Tag() quickfix.Tag { return tag.OrderQty }

// NewOrderQty returns a new OrderQtyField initialized with val and scale.
func NewOrderQty(val decimal.Decimal, scale int32) OrderQtyField {
	return OrderQtyField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f OrderQtyField) Value() (val decimal.Decimal) { return f.Decimal }

// OrigClOrdIDField is a STRING field.
type OrigClOrdIDField struct{ quickfix.FIXString }

// Tag returns tag.OrigClOrdID (41).
func (f OrigClOrdIDField) Tag() quickfix.Tag { return tag.OrigClOrdID }

// NewOrigClOrdID returns a new OrigClOrdIDField initialized with val.
func NewOrigClOrdID(val string) OrigClOrdIDField {
	return OrigClOrdIDField{quickfix.FIXString(val)}
}

func (f OrigClOrdIDField) Value() string { return f.String() }

// OrigSendingTimeField is a UTCTIMESTAMP field.
type OrigSendingTimeField struct{ quickfix.FIXUTCTimestamp }

// Tag returns tag.OrigSendingTime (122).
func (f OrigSendingTimeField) Tag() quickfix.Tag { return tag.OrigSendingTime }

// NewOrigSendingTime returns a new OrigSendingTimeField initialized with val.
func NewOrigSendingTime(val time.Time) OrigSendingTimeField {
	return NewOrigSendingTimeWithPrecision(val, quickfix.Millis)
}

// NewOrigSendingTimeNoMillis returns a new OrigSendingTimeField initialized with val without millisecs.
func NewOrigSendingTimeNoMillis(val time.Time) OrigSendingTimeField {
	return NewOrigSendingTimeWithPrecision(val, quickfix.Seconds)
}

// NewOrigSendingTimeWithPrecision returns a new OrigSendingTimeField initialized with val of specified precision.
func NewOrigSendingTimeWithPrecision(val time.Time, precision quickfix.TimestampPrecision) OrigSendingTimeField {
	return OrigSendingTimeField{quickfix.FIXUTCTimestamp{Time: val, Precision: precision}}
}

func (f OrigSendingTimeField) Value() time.Time { return f.Time }

// OrigTimeField is a UTCTIMESTAMP field.
type OrigTimeField struct{ quickfix.FIXUTCTimestamp }

// Tag returns tag.OrigTime (42).
func (f OrigTimeField) Tag() quickfix.Tag { return tag.OrigTime }

// NewOrigTime returns a new OrigTimeField initialized with val.
func NewOrigTime(val time.Time) OrigTimeField {
	return NewOrigTimeWithPrecision(val, quickfix.Millis)
}

// NewOrigTimeNoMillis returns a new OrigTimeField initialized with val without millisecs.
func NewOrigTimeNoMillis(val time.Time) OrigTimeField {
	return NewOrigTimeWithPrecision(val, quickfix.Seconds)
}

// NewOrigTimeWithPrecision returns a new OrigTimeField initialized with val of specified precision.
func NewOrigTimeWithPrecision(val time.Time, precision quickfix.TimestampPrecision) OrigTimeField {
	return OrigTimeField{quickfix.FIXUTCTimestamp{Time: val, Precision: precision}}
}

func (f OrigTimeField) Value() time.Time { return f.Time }

// PartyIDField is a STRING field.
type PartyIDField struct{ quickfix.FIXString }

// Tag returns tag.PartyID (448).
func (f PartyIDField) Tag() quickfix.Tag { return tag.PartyID }

// NewPartyID returns a new PartyIDField initialized with val.
func NewPartyID(val string) PartyIDField {
	return PartyIDField{quickfix.FIXString(val)}
}

func (f PartyIDField) Value() string { return f.String() }

// PartyIDSourceField is a enum.PartyIDSource field.
type PartyIDSourceField struct{ quickfix.FIXString }

// Tag returns tag.PartyIDSource (447).
func (f PartyIDSourceField) Tag() quickfix.Tag { return tag.PartyIDSource }

func NewPartyIDSource(val enum.PartyIDSource) PartyIDSourceField {
	return PartyIDSourceField{quickfix.FIXString(val)}
}

func (f PartyIDSourceField) Value() enum.PartyIDSource { return enum.PartyIDSource(f.String()) }

// PartyRoleField is a enum.PartyRole field.
type PartyRoleField struct{ quickfix.FIXString }

// Tag returns tag.PartyRole (452).
func (f PartyRoleField) Tag() quickfix.Tag { return tag.PartyRole }

func NewPartyRole(val enum.PartyRole) PartyRoleField {
	return PartyRoleField{quickfix.FIXString(val)}
}

func (f PartyRoleField) Value() enum.PartyRole { return enum.PartyRole(f.String()) }

// PartyRoleQualifierField is a enum.PartyRoleQualifier field.
type PartyRoleQualifierField struct{ quickfix.FIXString }

// Tag returns tag.PartyRoleQualifier (2376).
func (f PartyRoleQualifierField) Tag() quickfix.Tag { return tag.PartyRoleQualifier }

func NewPartyRoleQualifier(val enum.PartyRoleQualifier) PartyRoleQualifierField {
	return PartyRoleQualifierField{quickfix.FIXString(val)}
}

func (f PartyRoleQualifierField) Value() enum.PartyRoleQualifier {
	return enum.PartyRoleQualifier(f.String())
}

// PartySubIDField is a STRING field.
type PartySubIDField struct{ quickfix.FIXString }

// Tag returns tag.PartySubID (523).
func (f PartySubIDField) Tag() quickfix.Tag { return tag.PartySubID }

// NewPartySubID returns a new PartySubIDField initialized with val.
func NewPartySubID(val string) PartySubIDField {
	return PartySubIDField{quickfix.FIXString(val)}
}

func (f PartySubIDField) Value() string { return f.String() }

// PartySubIDTypeField is a enum.PartySubIDType field.
type PartySubIDTypeField struct{ quickfix.FIXString }

// Tag returns tag.PartySubIDType (803).
func (f PartySubIDTypeField) Tag() quickfix.Tag { return tag.PartySubIDType }

func NewPartySubIDType(val enum.PartySubIDType) PartySubIDTypeField {
	return PartySubIDTypeField{quickfix.FIXString(val)}
}

func (f PartySubIDTypeField) Value() enum.PartySubIDType { return enum.PartySubIDType(f.String()) }

// PasswordField is a STRING field.
type PasswordField struct{ quickfix.FIXString }

// Tag returns tag.Password (554).
func (f PasswordField) Tag() quickfix.Tag { return tag.Password }

// NewPassword returns a new PasswordField initialized with val.
func NewPassword(val string) PasswordField {
	return PasswordField{quickfix.FIXString(val)}
}

func (f PasswordField) Value() string { return f.String() }

// PossDupFlagField is a BOOLEAN field.
type PossDupFlagField struct{ quickfix.FIXBoolean }

// Tag returns tag.PossDupFlag (43).
func (f PossDupFlagField) Tag() quickfix.Tag { return tag.PossDupFlag }

// NewPossDupFlag returns a new PossDupFlagField initialized with val.
func NewPossDupFlag(val bool) PossDupFlagField {
	return PossDupFlagField{quickfix.FIXBoolean(val)}
}

func (f PossDupFlagField) Value() bool { return f.Bool() }

// PossResendField is a BOOLEAN field.
type PossResendField struct{ quickfix.FIXBoolean }

// Tag returns tag.PossResend (97).
func (f PossResendField) Tag() quickfix.Tag { return tag.PossResend }

// NewPossResend returns a new PossResendField initialized with val.
func NewPossResend(val bool) PossResendField {
	return PossResendField{quickfix.FIXBoolean(val)}
}

func (f PossResendField) Value() bool { return f.Bool() }

// PriceField is a PRICE field.
type PriceField struct{ quickfix.FIXDecimal }

// Tag returns tag.Price (44).
func (f PriceField) Tag() quickfix.Tag { return tag.Price }

// NewPrice returns a new PriceField initialized with val and scale.
func NewPrice(val decimal.Decimal, scale int32) PriceField {
	return PriceField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f PriceField) Value() (val decimal.Decimal) { return f.Decimal }

// QuantityField is a QTY field.
type QuantityField struct{ quickfix.FIXDecimal }

// Tag returns tag.Quantity (53).
func (f QuantityField) Tag() quickfix.Tag { return tag.Quantity }

// NewQuantity returns a new QuantityField initialized with val and scale.
func NewQuantity(val decimal.Decimal, scale int32) QuantityField {
	return QuantityField{quickfix.FIXDecimal{Decimal: val, Scale: scale}}
}

func (f QuantityField) Value() (val decimal.Decimal) { return f.Decimal }

// QuoteCancelTypeField is a enum.QuoteCancelType field.
type QuoteCancelTypeField struct{ quickfix.FIXString }

// Tag returns tag.QuoteCancelType (298).
func (f QuoteCancelTypeField) Tag() quickfix.Tag { return tag.QuoteCancelType }

func NewQuoteCancelType(val enum.QuoteCancelType) QuoteCancelTypeField {
	return QuoteCancelTypeField{quickfix.FIXString(val)}
}

func (f QuoteCancelTypeField) Value() enum.QuoteCancelType { return enum.QuoteCancelType(f.String()) }

// QuoteEntryIDField is a STRING field.
type QuoteEntryIDField struct{ quickfix.FIXString }

// Tag returns tag.QuoteEntryID (299).
func (f QuoteEntryIDField) Tag() quickfix.Tag { return tag.QuoteEntryID }

// NewQuoteEntryID returns a new QuoteEntryIDField initialized with val.
func NewQuoteEntryID(val string) QuoteEntryIDField {
	return QuoteEntryIDField{quickfix.FIXString(val)}
}

func (f QuoteEntryIDField) Value() string { return f.String() }

// QuoteIDField is a STRING field.
type QuoteIDField struct{ quickfix.FIXString }

// Tag returns tag.QuoteID (117).
func (f QuoteIDField) Tag() quickfix.Tag { return tag.QuoteID }

// NewQuoteID returns a new QuoteIDField initialized with val.
func NewQuoteID(val string) QuoteIDField {
	return QuoteIDField{quickfix.FIXString(val)}
}

func (f QuoteIDField) Value() string { return f.String() }

// QuoteMsgIDField is a STRING field.
type QuoteMsgIDField struct{ quickfix.FIXString }

// Tag returns tag.QuoteMsgID (1166).
func (f QuoteMsgIDField) Tag() quickfix.Tag { return tag.QuoteMsgID }

// NewQuoteMsgID returns a new QuoteMsgIDField initialized with val.
func NewQuoteMsgID(val string) QuoteMsgIDField {
	return QuoteMsgIDField{quickfix.FIXString(val)}
}

func (f QuoteMsgIDField) Value() string { return f.String() }

// QuoteRejectReasonField is a enum.QuoteRejectReason field.
type QuoteRejectReasonField struct{ quickfix.FIXString }

// Tag returns tag.QuoteRejectReason (300).
func (f QuoteRejectReasonField) Tag() quickfix.Tag { return tag.QuoteRejectReason }

func NewQuoteRejectReason(val enum.QuoteRejectReason) QuoteRejectReasonField {
	return QuoteRejectReasonField{quickfix.FIXString(val)}
}

func (f QuoteRejectReasonField) Value() enum.QuoteRejectReason {
	return enum.QuoteRejectReason(f.String())
}

// QuoteStatusField is a enum.QuoteStatus field.
type QuoteStatusField struct{ quickfix.FIXString }

// Tag returns tag.QuoteStatus (297).
func (f QuoteStatusField) Tag() quickfix.Tag { return tag.QuoteStatus }

func NewQuoteStatus(val enum.QuoteStatus) QuoteStatusField {
	return QuoteStatusField{quickfix.FIXString(val)}
}

func (f QuoteStatusField) Value() enum.QuoteStatus { return enum.QuoteStatus(f.String()) }

// RawDataField is a DATA field.
type RawDataField struct{ quickfix.FIXString }

// Tag returns tag.RawData (96).
func (f RawDataField) Tag() quickfix.Tag { return tag.RawData }

// NewRawData returns a new RawDataField initialized with val.
func NewRawData(val string) RawDataField {
	return RawDataField{quickfix.FIXString(val)}
}

func (f RawDataField) Value() string { return f.String() }

// RawDataLengthField is a LENGTH field.
type RawDataLengthField struct{ quickfix.FIXInt }

// Tag returns tag.RawDataLength (95).
func (f RawDataLengthField) Tag() quickfix.Tag { return tag.RawDataLength }

// NewRawDataLength returns a new RawDataLengthField initialized with val.
func NewRawDataLength(val int) RawDataLengthField {
	return RawDataLengthField{quickfix.FIXInt(val)}
}

func (f RawDataLengthField) Value() int { return f.Int() }

// RefApplVerIDField is a STRING field.
type RefApplVerIDField struct{ quickfix.FIXString }

// Tag returns tag.RefApplVerID (1130).
func (f RefApplVerIDField) Tag() quickfix.Tag { return tag.RefApplVerID }

// NewRefApplVerID returns a new RefApplVerIDField initialized with val.
func NewRefApplVerID(val string) RefApplVerIDField {
	return RefApplVerIDField{quickfix.FIXString(val)}
}

func (f RefApplVerIDField) Value() string { return f.String() }

// RefCstmApplVerIDField is a STRING field.
type RefCstmApplVerIDField struct{ quickfix.FIXString }

// Tag returns tag.RefCstmApplVerID (1131).
func (f RefCstmApplVerIDField) Tag() quickfix.Tag { return tag.RefCstmApplVerID }

// NewRefCstmApplVerID returns a new RefCstmApplVerIDField initialized with val.
func NewRefCstmApplVerID(val string) RefCstmApplVerIDField {
	return RefCstmApplVerIDField{quickfix.FIXString(val)}
}

func (f RefCstmApplVerIDField) Value() string { return f.String() }

// RefMsgTypeField is a STRING field.
type RefMsgTypeField struct{ quickfix.FIXString }

// Tag returns tag.RefMsgType (372).
func (f RefMsgTypeField) Tag() quickfix.Tag { return tag.RefMsgType }

// NewRefMsgType returns a new RefMsgTypeField initialized with val.
func NewRefMsgType(val string) RefMsgTypeField {
	return RefMsgTypeField{quickfix.FIXString(val)}
}

func (f RefMsgTypeField) Value() string { return f.String() }

// RefSeqNumField is a SEQNUM field.
type RefSeqNumField struct{ quickfix.FIXInt }

// Tag returns tag.RefSeqNum (45).
func (f RefSeqNumField) Tag() quickfix.Tag { return tag.RefSeqNum }

// NewRefSeqNum returns a new RefSeqNumField initialized with val.
func NewRefSeqNum(val int) RefSeqNumField {
	return RefSeqNumField{quickfix.FIXInt(val)}
}

func (f RefSeqNumField) Value() int { return f.Int() }

// RefTagIDField is a INT field.
type RefTagIDField struct{ quickfix.FIXInt }

// Tag returns tag.RefTagID (371).
func (f RefTagIDField) Tag() quickfix.Tag { return tag.RefTagID }

// NewRefTagID returns a new RefTagIDField initialized with val.
func NewRefTagID(val int) RefTagIDField {
	return RefTagIDField{quickfix.FIXInt(val)}
}

func (f RefTagIDField) Value() int { return f.Int() }

// ResetSeqNumFlagField is a BOOLEAN field.
type ResetSeqNumFlagField struct{ quickfix.FIXBoolean }

// Tag returns tag.ResetSeqNumFlag (141).
func (f ResetSeqNumFlagField) Tag() quickfix.Tag { return tag.ResetSeqNumFlag }

// NewResetSeqNumFlag returns a new ResetSeqNumFlagField initialized with val.
func NewResetSeqNumFlag(val bool) ResetSeqNumFlagField {
	return ResetSeqNumFlagField{quickfix.FIXBoolean(val)}
}

func (f ResetSeqNumFlagField) Value() bool { return f.Bool() }

// SecondaryOrderIDField is a STRING field.
type SecondaryOrderIDField struct{ quickfix.FIXString }

// Tag returns tag.SecondaryOrderID (198).
func (f SecondaryOrderIDField) Tag() quickfix.Tag { return tag.SecondaryOrderID }

// NewSecondaryOrderID returns a new SecondaryOrderIDField initialized with val.
func NewSecondaryOrderID(val string) SecondaryOrderIDField {
	return SecondaryOrderIDField{quickfix.FIXString(val)}
}

func (f SecondaryOrderIDField) Value() string { return f.String() }

// SecureDataField is a DATA field.
type SecureDataField struct{ quickfix.FIXString }

// Tag returns tag.SecureData (91).
func (f SecureDataField) Tag() quickfix.Tag { return tag.SecureData }

// NewSecureData returns a new SecureDataField initialized with val.
func NewSecureData(val string) SecureDataField {
	return SecureDataField{quickfix.FIXString(val)}
}

func (f SecureDataField) Value() string { return f.String() }

// SecureDataLenField is a LENGTH field.
type SecureDataLenField struct{ quickfix.FIXInt }

// Tag returns tag.SecureDataLen (90).
func (f SecureDataLenField) Tag() quickfix.Tag { return tag.SecureDataLen }

// NewSecureDataLen returns a new SecureDataLenField initialized with val.
func NewSecureDataLen(val int) SecureDataLenField {
	return SecureDataLenField{quickfix.FIXInt(val)}
}

func (f SecureDataLenField) Value() int { return f.Int() }

// SecurityIDField is a STRING field.
type SecurityIDField struct{ quickfix.FIXString }

// Tag returns tag.SecurityID (48).
func (f SecurityIDField) Tag() quickfix.Tag { return tag.SecurityID }

// NewSecurityID returns a new SecurityIDField initialized with val.
func NewSecurityID(val string) SecurityIDField {
	return SecurityIDField{quickfix.FIXString(val)}
}

func (f SecurityIDField) Value() string { return f.String() }

// SecurityIDSourceField is a enum.SecurityIDSource field.
type SecurityIDSourceField struct{ quickfix.FIXString }

// Tag returns tag.SecurityIDSource (22).
func (f SecurityIDSourceField) Tag() quickfix.Tag { return tag.SecurityIDSource }

func NewSecurityIDSource(val enum.SecurityIDSource) SecurityIDSourceField {
	return SecurityIDSourceField{quickfix.FIXString(val)}
}

func (f SecurityIDSourceField) Value() enum.SecurityIDSource {
	return enum.SecurityIDSource(f.String())
}

// SecurityListRequestTypeField is a enum.SecurityListRequestType field.
type SecurityListRequestTypeField struct{ quickfix.FIXString }

// Tag returns tag.SecurityListRequestType (559).
func (f SecurityListRequestTypeField) Tag() quickfix.Tag { return tag.SecurityListRequestType }

func NewSecurityListRequestType(val enum.SecurityListRequestType) SecurityListRequestTypeField {
	return SecurityListRequestTypeField{quickfix.FIXString(val)}
}

func (f SecurityListRequestTypeField) Value() enum.SecurityListRequestType {
	return enum.SecurityListRequestType(f.String())
}

// SecurityReqIDField is a STRING field.
type SecurityReqIDField struct{ quickfix.FIXString }

// Tag returns tag.SecurityReqID (320).
func (f SecurityReqIDField) Tag() quickfix.Tag { return tag.SecurityReqID }

// NewSecurityReqID returns a new SecurityReqIDField initialized with val.
func NewSecurityReqID(val string) SecurityReqIDField {
	return SecurityReqIDField{quickfix.FIXString(val)}
}

func (f SecurityReqIDField) Value() string { return f.String() }

// SecurityRequestResultField is a enum.SecurityRequestResult field.
type SecurityRequestResultField struct{ quickfix.FIXString }

// Tag returns tag.SecurityRequestResult (560).
func (f SecurityRequestResultField) Tag() quickfix.Tag { return tag.SecurityRequestResult }

func NewSecurityRequestResult(val enum.SecurityRequestResult) SecurityRequestResultField {
	return SecurityRequestResultField{quickfix.FIXString(val)}
}

func (f SecurityRequestResultField) Value() enum.SecurityRequestResult {
	return enum.SecurityRequestResult(f.String())
}

// SecurityRequestTypeField is a enum.SecurityRequestType field.
type SecurityRequestTypeField struct{ quickfix.FIXString }

// Tag returns tag.SecurityRequestType (321).
func (f SecurityRequestTypeField) Tag() quickfix.Tag { return tag.SecurityRequestType }

func NewSecurityRequestType(val enum.SecurityRequestType) SecurityRequestTypeField {
	return SecurityRequestTypeField{quickfix.FIXString(val)}
}

func (f SecurityRequestTypeField) Value() enum.SecurityRequestType {
	return enum.SecurityRequestType(f.String())
}

// SecurityResponseIDField is a STRING field.
type SecurityResponseIDField struct{ quickfix.FIXString }

// Tag returns tag.SecurityResponseID (322).
func (f SecurityResponseIDField) Tag() quickfix.Tag { return tag.SecurityResponseID }

// NewSecurityResponseID returns a new SecurityResponseIDField initialized with val.
func NewSecurityResponseID(val string) SecurityResponseIDField {
	return SecurityResponseIDField{quickfix.FIXString(val)}
}

func (f SecurityResponseIDField) Value() string { return f.String() }

// SecurityStatusReqIDField is a STRING field.
type SecurityStatusReqIDField struct{ quickfix.FIXString }

// Tag returns tag.SecurityStatusReqID (324).
func (f SecurityStatusReqIDField) Tag() quickfix.Tag { return tag.SecurityStatusReqID }

// NewSecurityStatusReqID returns a new SecurityStatusReqIDField initialized with val.
func NewSecurityStatusReqID(val string) SecurityStatusReqIDField {
	return SecurityStatusReqIDField{quickfix.FIXString(val)}
}

func (f SecurityStatusReqIDField) Value() string { return f.String() }

// SecurityTradingStatusField is a enum.SecurityTradingStatus field.
type SecurityTradingStatusField struct{ quickfix.FIXString }

// Tag returns tag.SecurityTradingStatus (326).
func (f SecurityTradingStatusField) Tag() quickfix.Tag { return tag.SecurityTradingStatus }

func NewSecurityTradingStatus(val enum.SecurityTradingStatus) SecurityTradingStatusField {
	return SecurityTradingStatusField{quickfix.FIXString(val)}
}

func (f SecurityTradingStatusField) Value() enum.SecurityTradingStatus {
	return enum.SecurityTradingStatus(f.String())
}

// SenderCompIDField is a STRING field.
type SenderCompIDField struct{ quickfix.FIXString }

// Tag returns tag.SenderCompID (49).
func (f SenderCompIDField) Tag() quickfix.Tag { return tag.SenderCompID }

// NewSenderCompID returns a new SenderCompIDField initialized with val.
func NewSenderCompID(val string) SenderCompIDField {
	return SenderCompIDField{quickfix.FIXString(val)}
}

func (f SenderCompIDField) Value() string { return f.String() }

// SenderLocationIDField is a STRING field.
type SenderLocationIDField struct{ quickfix.FIXString }

// Tag returns tag.SenderLocationID (142).
func (f SenderLocationIDField) Tag() quickfix.Tag { return tag.SenderLocationID }

// NewSenderLocationID returns a new SenderLocationIDField initialized with val.
func NewSenderLocationID(val string) SenderLocationIDField {
	return SenderLocationIDField{quickfix.FIXString(val)}
}

func (f SenderLocationIDField) Value() string { return f.String() }

// SenderSubIDField is a STRING field.
type SenderSubIDField struct{ quickfix.FIXString }

// Tag returns tag.SenderSubID (50).
func (f SenderSubIDField) Tag() quickfix.Tag { return tag.SenderSubID }

// NewSenderSubID returns a new SenderSubIDField initialized with val.
func NewSenderSubID(val string) SenderSubIDField {
	return SenderSubIDField{quickfix.FIXString(val)}
}

func (f SenderSubIDField) Value() string { return f.String() }

// SendingDateField is a LOCALMKTDATE field.
type SendingDateField struct{ quickfix.FIXString }

// Tag returns tag.SendingDate (51).
func (f SendingDateField) Tag() quickfix.Tag { return tag.SendingDate }

// NewSendingDate returns a new SendingDateField initialized with val.
func NewSendingDate(val string) SendingDateField {
	return SendingDateField{quickfix.FIXString(val)}
}

func (f SendingDateField) Value() string { return f.String() }

// SendingTimeField is a UTCTIMESTAMP field.
type SendingTimeField struct{ quickfix.FIXUTCTimestamp }

// Tag returns tag.SendingTime (52).
func (f SendingTimeField) Tag() quickfix.Tag { return tag.SendingTime }

// NewSendingTime returns a new SendingTimeField initialized with val.
func NewSendingTime(val time.Time) SendingTimeField {
	return NewSendingTimeWithPrecision(val, quickfix.Millis)
}

// NewSendingTimeNoMillis returns a new SendingTimeField initialized with val without millisecs.
func NewSendingTimeNoMillis(val time.Time) SendingTimeField {
	return NewSendingTimeWithPrecision(val, quickfix.Seconds)
}

// NewSendingTimeWithPrecision returns a new SendingTimeField initialized with val of specified precision.
func NewSendingTimeWithPrecision(val time.Time, precision quickfix.TimestampPrecision) SendingTimeField {
	return SendingTimeField{quickfix.FIXUTCTimestamp{Time: val, Precision: precision}}
}

func (f SendingTimeField) Value() time.Time { return f.Time }

// SessionRejectReasonField is a enum.SessionRejectReason field.
type SessionRejectReasonField struct{ quickfix.FIXString }

// Tag returns tag.SessionRejectReason (373).
func (f SessionRejectReasonField) Tag() quickfix.Tag { return tag.SessionRejectReason }

func NewSessionRejectReason(val enum.SessionRejectReason) SessionRejectReasonField {
	return SessionRejectReasonField{quickfix.FIXString(val)}
}

func (f SessionRejectReasonField) Value() enum.SessionRejectReason {
	return enum.SessionRejectReason(f.String())
}

// SideField is a enum.Side field.
type SideField struct{ quickfix.FIXString }

// Tag returns tag.Side (54).
func (f SideField) Tag() quickfix.Tag { return tag.Side }

func NewSide(val enum.Side) SideField {
	return SideField{quickfix.FIXString(val)}
}

func (f SideField) Value() enum.Side { return enum.Side(f.String()) }

// SignatureField is a DATA field.
type SignatureField struct{ quickfix.FIXString }

// Tag returns tag.Signature (89).
func (f SignatureField) Tag() quickfix.Tag { return tag.Signature }

// NewSignature returns a new SignatureField initialized with val.
func NewSignature(val string) SignatureField {
	return SignatureField{quickfix.FIXString(val)}
}

func (f SignatureField) Value() string { return f.String() }

// SignatureLengthField is a LENGTH field.
type SignatureLengthField struct{ quickfix.FIXInt }

// Tag returns tag.SignatureLength (93).
func (f SignatureLengthField) Tag() quickfix.Tag { return tag.SignatureLength }

// NewSignatureLength returns a new SignatureLengthField initialized with val.
func NewSignatureLength(val int) SignatureLengthField {
	return SignatureLengthField{quickfix.FIXInt(val)}
}

func (f SignatureLengthField) Value() int { return f.Int() }

// SubscriptionRequestTypeField is a enum.SubscriptionRequestType field.
type SubscriptionRequestTypeField struct{ quickfix.FIXString }

// Tag returns tag.SubscriptionRequestType (263).
func (f SubscriptionRequestTypeField) Tag() quickfix.Tag { return tag.SubscriptionRequestType }

func NewSubscriptionRequestType(val enum.SubscriptionRequestType) SubscriptionRequestTypeField {
	return SubscriptionRequestTypeField{quickfix.FIXString(val)}
}

func (f SubscriptionRequestTypeField) Value() enum.SubscriptionRequestType {
	return enum.SubscriptionRequestType(f.String())
}

// SymbolField is a STRING field.
type SymbolField struct{ quickfix.FIXString }

// Tag returns tag.Symbol (55).
func (f SymbolField) Tag() quickfix.Tag { return tag.Symbol }

// NewSymbol returns a new SymbolField initialized with val.
func NewSymbol(val string) SymbolField {
	return SymbolField{quickfix.FIXString(val)}
}

func (f SymbolField) Value() string { return f.String() }

// TargetCompIDField is a STRING field.
type TargetCompIDField struct{ quickfix.FIXString }

// Tag returns tag.TargetCompID (56).
func (f TargetCompIDField) Tag() quickfix.Tag { return tag.TargetCompID }

// NewTargetCompID returns a new TargetCompIDField initialized with val.
func NewTargetCompID(val string) TargetCompIDField {
	return TargetCompIDField{quickfix.FIXString(val)}
}

func (f TargetCompIDField) Value() string { return f.String() }

// TargetLocationIDField is a STRING field.
type TargetLocationIDField struct{ quickfix.FIXString }

// Tag returns tag.TargetLocationID (143).
func (f TargetLocationIDField) Tag() quickfix.Tag { return tag.TargetLocationID }

// NewTargetLocationID returns a new TargetLocationIDField initialized with val.
func NewTargetLocationID(val string) TargetLocationIDField {
	return TargetLocationIDField{quickfix.FIXString(val)}
}

func (f TargetLocationIDField) Value() string { return f.String() }

// TargetSubIDField is a STRING field.
type TargetSubIDField struct{ quickfix.FIXString }

// Tag returns tag.TargetSubID (57).
func (f TargetSubIDField) Tag() quickfix.Tag { return tag.TargetSubID }

// NewTargetSubID returns a new TargetSubIDField initialized with val.
func NewTargetSubID(val string) TargetSubIDField {
	return TargetSubIDField{quickfix.FIXString(val)}
}

func (f TargetSubIDField) Value() string { return f.String() }

// TestMessageIndicatorField is a BOOLEAN field.
type TestMessageIndicatorField struct{ quickfix.FIXBoolean }

// Tag returns tag.TestMessageIndicator (464).
func (f TestMessageIndicatorField) Tag() quickfix.Tag { return tag.TestMessageIndicator }

// NewTestMessageIndicator returns a new TestMessageIndicatorField initialized with val.
func NewTestMessageIndicator(val bool) TestMessageIndicatorField {
	return TestMessageIndicatorField{quickfix.FIXBoolean(val)}
}

func (f TestMessageIndicatorField) Value() bool { return f.Bool() }

// TestReqIDField is a STRING field.
type TestReqIDField struct{ quickfix.FIXString }

// Tag returns tag.TestReqID (112).
func (f TestReqIDField) Tag() quickfix.Tag { return tag.TestReqID }

// NewTestReqID returns a new TestReqIDField initialized with val.
func NewTestReqID(val string) TestReqIDField {
	return TestReqIDField{quickfix.FIXString(val)}
}

func (f TestReqIDField) Value() string { return f.String() }

// TextField is a STRING field.
type TextField struct{ quickfix.FIXString }

// Tag returns tag.Text (58).
func (f TextField) Tag() quickfix.Tag { return tag.Text }

// NewText returns a new TextField initialized with val.
func NewText(val string) TextField {
	return TextField{quickfix.FIXString(val)}
}

func (f TextField) Value() string { return f.String() }

// TimeInForceField is a enum.TimeInForce field.
type TimeInForceField struct{ quickfix.FIXString }

// Tag returns tag.TimeInForce (59).
func (f TimeInForceField) Tag() quickfix.Tag { return tag.TimeInForce }

func NewTimeInForce(val enum.TimeInForce) TimeInForceField {
	return TimeInForceField{quickfix.FIXString(val)}
}

func (f TimeInForceField) Value() enum.TimeInForce { return enum.TimeInForce(f.String()) }

// TotNoRelatedSymField is a INT field.
type TotNoRelatedSymField struct{ quickfix.FIXInt }

// Tag returns tag.TotNoRelatedSym (393).
func (f TotNoRelatedSymField) Tag() quickfix.Tag { return tag.TotNoRelatedSym }

// NewTotNoRelatedSym returns a new TotNoRelatedSymField initialized with val.
func NewTotNoRelatedSym(val int) TotNoRelatedSymField {
	return TotNoRelatedSymField{quickfix.FIXInt(val)}
}

func (f TotNoRelatedSymField) Value() int { return f.Int() }

// TotNumReportsField is a INT field.
type TotNumReportsField struct{ quickfix.FIXInt }

// Tag returns tag.TotNumReports (911).
func (f TotNumReportsField) Tag() quickfix.Tag { return tag.TotNumReports }

// NewTotNumReports returns a new TotNumReportsField initialized with val.
func NewTotNumReports(val int) TotNumReportsField {
	return TotNumReportsField{quickfix.FIXInt(val)}
}

func (f TotNumReportsField) Value() int { return f.Int() }

// TradeConditionField is a enum.TradeCondition field.
type TradeConditionField struct{ quickfix.FIXString }

// Tag returns tag.TradeCondition (277).
func (f TradeConditionField) Tag() quickfix.Tag { return tag.TradeCondition }

func NewTradeCondition(val enum.TradeCondition) TradeConditionField {
	return TradeConditionField{quickfix.FIXString(val)}
}

func (f TradeConditionField) Value() enum.TradeCondition { return enum.TradeCondition(f.String()) }

// TradeIDField is a STRING field.
type TradeIDField struct{ quickfix.FIXString }

// Tag returns tag.TradeID (1003).
func (f TradeIDField) Tag() quickfix.Tag { return tag.TradeID }

// NewTradeID returns a new TradeIDField initialized with val.
func NewTradeID(val string) TradeIDField {
	return TradeIDField{quickfix.FIXString(val)}
}

func (f TradeIDField) Value() string { return f.String() }

// TradingSessionIDField is a enum.TradingSessionID field.
type TradingSessionIDField struct{ quickfix.FIXString }

// Tag returns tag.TradingSessionID (336).
func (f TradingSessionIDField) Tag() quickfix.Tag { return tag.TradingSessionID }

func NewTradingSessionID(val enum.TradingSessionID) TradingSessionIDField {
	return TradingSessionIDField{quickfix.FIXString(val)}
}

func (f TradingSessionIDField) Value() enum.TradingSessionID {
	return enum.TradingSessionID(f.String())
}

// TransactTimeField is a UTCTIMESTAMP field.
type TransactTimeField struct{ quickfix.FIXUTCTimestamp }

// Tag returns tag.TransactTime (60).
func (f TransactTimeField) Tag() quickfix.Tag { return tag.TransactTime }

// NewTransactTime returns a new TransactTimeField initialized with val.
func NewTransactTime(val time.Time) TransactTimeField {
	return NewTransactTimeWithPrecision(val, quickfix.Millis)
}

// NewTransactTimeNoMillis returns a new TransactTimeField initialized with val without millisecs.
func NewTransactTimeNoMillis(val time.Time) TransactTimeField {
	return NewTransactTimeWithPrecision(val, quickfix.Seconds)
}

// NewTransactTimeWithPrecision returns a new TransactTimeField initialized with val of specified precision.
func NewTransactTimeWithPrecision(val time.Time, precision quickfix.TimestampPrecision) TransactTimeField {
	return TransactTimeField{quickfix.FIXUTCTimestamp{Time: val, Precision: precision}}
}

func (f TransactTimeField) Value() time.Time { return f.Time }

// TrdMatchIDField is a STRING field.
type TrdMatchIDField struct{ quickfix.FIXString }

// Tag returns tag.TrdMatchID (880).
func (f TrdMatchIDField) Tag() quickfix.Tag { return tag.TrdMatchID }

// NewTrdMatchID returns a new TrdMatchIDField initialized with val.
func NewTrdMatchID(val string) TrdMatchIDField {
	return TrdMatchIDField{quickfix.FIXString(val)}
}

func (f TrdMatchIDField) Value() string { return f.String() }

// URLLinkField is a STRING field.
type URLLinkField struct{ quickfix.FIXString }

// Tag returns tag.URLLink (149).
func (f URLLinkField) Tag() quickfix.Tag { return tag.URLLink }

// NewURLLink returns a new URLLinkField initialized with val.
func NewURLLink(val string) URLLinkField {
	return URLLinkField{quickfix.FIXString(val)}
}

func (f URLLinkField) Value() string { return f.String() }

// UrgencyField is a enum.Urgency field.
type UrgencyField struct{ quickfix.FIXString }

// Tag returns tag.Urgency (61).
func (f UrgencyField) Tag() quickfix.Tag { return tag.Urgency }

func NewUrgency(val enum.Urgency) UrgencyField {
	return UrgencyField{quickfix.FIXString(val)}
}

func (f UrgencyField) Value() enum.Urgency { return enum.Urgency(f.String()) }

// UsernameField is a STRING field.
type UsernameField struct{ quickfix.FIXString }

// Tag returns tag.Username (553).
func (f UsernameField) Tag() quickfix.Tag { return tag.Username }

// NewUsername returns a new UsernameField initialized with val.
func NewUsername(val string) UsernameField {
	return UsernameField{quickfix.FIXString(val)}
}

func (f UsernameField) Value() string { return f.String() }

// XmlDataField is a DATA field.
type XmlDataField struct{ quickfix.FIXString }

// Tag returns tag.XmlData (213).
func (f XmlDataField) Tag() quickfix.Tag { return tag.XmlData }

// NewXmlData returns a new XmlDataField initialized with val.
func NewXmlData(val string) XmlDataField {
	return XmlDataField{quickfix.FIXString(val)}
}

func (f XmlDataField) Value() string { return f.String() }

// XmlDataLenField is a LENGTH field.
type XmlDataLenField struct{ quickfix.FIXInt }

// Tag returns tag.XmlDataLen (212).
func (f XmlDataLenField) Tag() quickfix.Tag { return tag.XmlDataLen }

// NewXmlDataLen returns a new XmlDataLenField initialized with val.
func NewXmlDataLen(val int) XmlDataLenField {
	return XmlDataLenField{quickfix.FIXInt(val)}
}

func (f XmlDataLenField) Value() int { return f.Int() }
