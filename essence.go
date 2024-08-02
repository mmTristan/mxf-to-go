// Copyright ©2019-2024  Mr MXF   info@mrmxf.com
// BSD-3-Clause License  https://opensource.org/license/bsd-3-clause/
package mxf2go

type essenceInformation struct {
	UL               string `xml:"UL,omitempty"`
	Name             string `xml:"Name,omitempty"`
	Symbol           string `xml:"Symbol,omitempty"`
	Definition       string `xml:"Definition,omitempty"`
	DefiningDocument string `xml:"DefiningDocument,omitempty"`
	IsDeprecated     bool   `xml:"IsDeprecated,omitempty"`
}

var EssenceLookUp = map[string]essenceInformation{
	"urn:smpte:ul:060e2b34.01020101.0d010301.0501017f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.0501017f", Name: "Type D-10 Element", Symbol: "TypeD10Element", Definition: "Identifies a Type D-10 constrained MPEG2 4:2:2 coded element (see SMPTE 331)", DefiningDocument: "SMPTE ST 386", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.0601107f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.0601107f", Name: "8-Ch AES3 Element", Symbol: "_8ChAES3Element", Definition: "Identifies a 8 channel AES3 audio data element", DefiningDocument: "SMPTE ST 386", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.077f207f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.077f207f", Name: "VBI Line Element", Symbol: "VBILineElement", Definition: "Identifies a VBI line data element", DefiningDocument: "SMPTE ST 331", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.077f217f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.077f217f", Name: "ANC Packet Element", Symbol: "ANCPacketElement", Definition: "Identifies a Ancillary Packet data element", DefiningDocument: "SMPTE ST 331", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.077f227f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.077f227f", Name: "General Data Element", Symbol: "GeneralDataElement", Definition: "Identifies a General data element", DefiningDocument: "SMPTE ST 331", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.077f407f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.077f407f", Name: "BWF Element", Symbol: "BWFElement", Definition: "Identifies a VBI line data element", DefiningDocument: "SMPTE ST 331", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.077f417f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.077f417f", Name: "JFIF Element", Symbol: "JFIFElement", Definition: "Identifies a JFIF data element", DefiningDocument: "SMPTE ST 331", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.077f427f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.077f427f", Name: "TIFF Element", Symbol: "TIFFElement", Definition: "Identifies a TIFF data element", DefiningDocument: "SMPTE ST 331", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.077f787f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.077f787f", Name: "Control Element", Symbol: "ControlElement", Definition: "Identifies a Control data element", DefiningDocument: "SMPTE ST 331", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.15010101": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.15010101", Name: "Type D-11 Element (Frame-Wrapped)", Symbol: "TypeD11ElementFrameWrapped", Definition: "Identifies a frame-wrapped Type D-11 picture element", DefiningDocument: "SMPTE ST 387", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.157f027f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.157f027f", Name: "Frame-wrapped Picture Element", Symbol: "FrameWrappedPictureElement", Definition: "Identifies a frame-wrapped uncompressed picture element", DefiningDocument: "SMPTE ST 384", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.157f037f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.157f037f", Name: "Clip-wrapped Picture Element", Symbol: "ClipWrappedPictureElement", Definition: "Identifies a clip-wrapped uncompressed picture element", DefiningDocument: "SMPTE ST 384", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.157f047f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.157f047f", Name: "Line-wrapped Picture Element", Symbol: "LineWrappedPictureElement", Definition: "Identifies a line-wrapped uncompressed picture element", DefiningDocument: "SMPTE ST 384", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.157f057f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.157f057f", Name: "Frame-wrapped MPEG Picture Element", Symbol: "FrameWrappedMPEGPictureElement", Definition: "Identifies a frame-wrapped MPEG picture element", DefiningDocument: "SMPTE ST 381", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.157f067f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.157f067f", Name: "Clip-wrapped MPEG Picture Element", Symbol: "ClipWrappedMPEGPictureElement", Definition: "Identifies a clip-wrapped MPEG picture element", DefiningDocument: "SMPTE ST 381", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.157f077f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.157f077f", Name: "Custom-wrapped MPEG Picture Element", Symbol: "CustomWrappedMPEGPictureElement", Definition: "Identifies a custom-wrapped MPEG picture element", DefiningDocument: "SMPTE ST 381", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.157f087f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.157f087f", Name: "Frame-wrapped JPEG 2000 Picture Element", Symbol: "FrameWrappedJPEG2000PictureElement", Definition: "Identifies a frame-wrapped JPEG 2000 Picture Element", DefiningDocument: "SMPTE ST 422", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.157f097f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.157f097f", Name: "Clip-wrapped JPEG 2000 Picture Element", Symbol: "ClipWrappedJPEG2000PictureElement", Definition: "Identifies a clip-wrapped JPEG 2000 Picture Element", DefiningDocument: "SMPTE ST 422", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.157f0a7f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.157f0a7f", Name: "Frame-wrapped VC-1 Picture Element", Symbol: "FrameWrappedVC1PictureElement", Definition: "Identifies a frame-wrapped VC-1 Picture Element", DefiningDocument: "SMPTE ST 2037", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.157f0b7f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.157f0b7f", Name: "Clip-wrapped VC-1 Picture Element", Symbol: "ClipWrappedVC1PictureElement", Definition: "Identifies a clip-wrapped VC-1 Picture Element", DefiningDocument: "SMPTE ST 2037", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.157f0c7f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.157f0c7f", Name: "Frame-wrapped VC-3 Picture Element", Symbol: "FrameWrappedVC3PictureElement", Definition: "Identifies a frame-wrapped VC-3 Picture Element", DefiningDocument: "SMPTE ST 2019-4", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.157f0d7f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.157f0d7f", Name: "Clip-wrapped VC-3 Picture Element", Symbol: "ClipWrappedVC3PictureElement", Definition: "Identifies a clip-wrapped VC-3 Picture Element", DefiningDocument: "SMPTE ST 2019-4", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.157f0e7f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.157f0e7f", Name: "Frame-wrapped TIFF/EP Picture Element", Symbol: "FrameWrappedTIFF_EPPictureElement", Definition: "Identifies a frame-wrapped TIFF/EP Picture Element", DefiningDocument: "SMPTE ST 2055", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.157f0f7f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.157f0f7f", Name: "Clip-wrapped TIFF/EP Picture Element", Symbol: "ClipWrappedTIFF_EPPictureElement", Definition: "Identifies a clip-wrapped TIFF/EP Picture Element", DefiningDocument: "SMPTE ST 2055", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.157f107f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.157f107f", Name: "Frame-wrapped VC-2 Picture Element", Symbol: "FrameWrappedVC2PictureElement", Definition: "Identifies a frame-wrapped VC-2 Picture Element", DefiningDocument: "SMPTE ST 2042-4", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.157f117f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.157f117f", Name: "Clip-wrapped VC-2 Picture Element", Symbol: "ClipWrappedVC2PictureElement", Definition: "Identifies a clip-wrapped VC-2 Picture Element", DefiningDocument: "SMPTE ST 2042-4", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.157f127f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.157f127f", Name: "Frame-wrapped ACES Picture Element", Symbol: "FrameWrappedACESPictureElement", Definition: "Identifies a frame-wrapped ACES Picture Element", DefiningDocument: "SMPTE ST 2065-5", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.157f137f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.157f137f", Name: "Clip-wrapped ACES Picture Element", Symbol: "ClipWrappedACESPictureElement", Definition: "Identifies a clip-wrapped ACES Picture Element", DefiningDocument: "SMPTE ST 2065-5", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.157f147f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.157f147f", Name: "Frame-wrapped VC-5 Picture Element", Symbol: "FrameWrappedVC5PictureElement", Definition: "Identifies a frame-wrapped VC-5 Picture Element", DefiningDocument: "SMPTE ST 2073-10", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.157f157f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.157f157f", Name: "Clip-wrapped VC-5 Picture Element", Symbol: "ClipWrappedVC5PictureElement", Definition: "Identifies a clip-wrapped VC-5 Picture Element", DefiningDocument: "SMPTE ST 2073-10", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.157f167f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.157f167f", Name: "Custom-wrapped VC-5 Picture Element", Symbol: "CustomWrappedVC5PictureElement", Definition: "Identifies a custom-wrapped VC-5 Picture Element", DefiningDocument: "SMPTE ST 2073-10", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.157f177f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.157f177f", Name: "Frame-wrapped ProRes Picture Element", Symbol: "FrameWrappedProResPictureElement", Definition: "Identifies a frame-wrapped ProRes Picture Element ", DefiningDocument: "SMPTE RDD 44", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.157f187f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.157f187f", Name: "Frame-wrapped DNxPacked Picture Element", Symbol: "FrameWrappedDNxPackedPictureElement", Definition: "Identifies Frame-wrapped DNxPacked Picture Element", DefiningDocument: "SMPTE RDD 50", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.157f197f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.157f197f", Name: "Clip-wrapped DNxPacked Picture Element", Symbol: "ClipWrappedDNxPackedPictureElement", Definition: "Identifies Clip-wrapped DNxPacked Picture Element", DefiningDocument: "SMPTE RDD 50", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.167f017f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.167f017f", Name: "Wave Frame-wrapped Sound Element", Symbol: "WaveFrameWrappedSoundElement", Definition: "Identifies a frame-wrapped Broadcast Wave sound element", DefiningDocument: "SMPTE ST 382", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.167f027f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.167f027f", Name: "Wave Clip-wrapped Sound Element", Symbol: "WaveClipWrappedSoundElement", Definition: "Identifies a clip-wrapped Broadcast Wave sound element", DefiningDocument: "SMPTE ST 382", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.167f037f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.167f037f", Name: "AES3 Frame-wrapped Sound Element", Symbol: "AES3FrameWrappedSoundElement", Definition: "Identifies a frame-wrapped AES3 sound element", DefiningDocument: "SMPTE ST 382", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.167f047f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.167f047f", Name: "AES3 Clip-wrapped Sound Element", Symbol: "AES3ClipWrappedSoundElement", Definition: "Identifies a clip-wrapped AES3 sound element", DefiningDocument: "SMPTE ST 382", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.167f057f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.167f057f", Name: "Frame-wrapped MPEG Sound Element", Symbol: "FrameWrappedMPEGSoundElement", Definition: "Identifies a frame-wrapped MPEG coded sound element", DefiningDocument: "SMPTE ST 381", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.167f067f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.167f067f", Name: "Clip-wrapped MPEG Sound Element", Symbol: "ClipWrappedMPEGSoundElement", Definition: "Identifies a clip-wrapped MPEG coded sound element", DefiningDocument: "SMPTE ST 381", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.167f077f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.167f077f", Name: "Custom-wrapped MPEG Sound Element", Symbol: "CustomWrappedMPEGSoundElement", Definition: "Identifies a custom-wrapped MPEG coded sound element", DefiningDocument: "SMPTE ST 381", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.167f087f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.167f087f", Name: "Frame-wrapped A-law Sound Element", Symbol: "FrameWrappedAlawSoundElement", Definition: "Identifies a frame-wrapped A-law coded sound element", DefiningDocument: "SMPTE ST 388", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.167f097f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.167f097f", Name: "Clip-wrapped A-law Sound Element", Symbol: "ClipWrappedAlawSoundElement", Definition: "Identifies a clip-wrapped A-law coded sound element", DefiningDocument: "SMPTE ST 388", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.167f0a7f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.167f0a7f", Name: "Custom-wrapped A-law Sound Element", Symbol: "CustomWrappedAlawSoundElement", Definition: "Identifies a custom-wrapped A-law coded sound element", DefiningDocument: "SMPTE ST 388", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.167f0b7f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.167f0b7f", Name: "Wave Custom-wrapped Sound Element", Symbol: "WaveCustomWrappedSoundElement", Definition: "Identifies a custom-wrapped Wav coded sound element", DefiningDocument: "SMPTE ST 382", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.167f0c7f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.167f0c7f", Name: "AES3 Custom-wrapped Sound Element", Symbol: "AES3CustomWrappedSoundElement", Definition: "Identifies a custom-wrapped AES3 coded sound element", DefiningDocument: "SMPTE ST 382", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.167f0d7f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.167f0d7f", Name: "IMF IAB Essence Clip-Wrapped Element", Symbol: "IMF_IABEssenceClipWrappedElement", Definition: "Identifies a clip-wrapped Immersive Audio Bitstream sound element", DefiningDocument: "SMPTE ST 2067-201", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.177f0101": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.177f0101", Name: "Frame-Wrapped VBI Data Element", Symbol: "FrameWrappedVBIDataElement", Definition: "Identifies a frame-Wrapped VBI Data Element", DefiningDocument: "SMPTE ST 436", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.177f0201": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.177f0201", Name: "Frame-Wrapped ANC Data Element", Symbol: "FrameWrappedANCDataElement", Definition: "Identifies a frame-Wrapped ANC Data Element", DefiningDocument: "SMPTE ST 436", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.177f057f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.177f057f", Name: "Frame-wrapped MPEG Data Element", Symbol: "FrameWrappedMPEGDataElement", Definition: "Identifies a frame-wrapped MPEG data element", DefiningDocument: "SMPTE ST 381", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.177f067f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.177f067f", Name: "Clip-wrapped MPEG Data Element", Symbol: "ClipWrappedMPEGDataElement", Definition: "Identifies a clip-wrapped MPEG data element", DefiningDocument: "SMPTE ST 381", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.177f077f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.177f077f", Name: "Custom-wrapped MPEG Data Element", Symbol: "CustomWrappedMPEGDataElement", Definition: "Identifies a custom-wrapped MPEG data element", DefiningDocument: "SMPTE ST 381", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.177f087f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.177f087f", Name: "Line-wrapped Picture Data Element", Symbol: "LineWrappedPictureDataElement", Definition: "Identifies a line-wrapped picture data element", DefiningDocument: "SMPTE ST 384", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.177f097f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.177f097f", Name: "Line-wrapped Picture V-ANC Data Element", Symbol: "LineWrappedPictureVANCDataElement", Definition: "Identifies a line-wrapped picture V-Anc packet data element", DefiningDocument: "SMPTE ST 384", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.177f0a7f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.177f0a7f", Name: "Line-wrapped Picture H-ANC Data Element", Symbol: "LineWrappedPictureHANCDataElement", Definition: "Identifies a line-wrapped picture H-Anc packet data element", DefiningDocument: "SMPTE ST 384", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.177f0b7f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.177f0b7f", Name: "Timed Text Data Element", Symbol: "TimedTextDataElement", Definition: "Identifies a timed text data element", DefiningDocument: "SMPTE ST 429-5", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.177f0d7f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.177f0d7f", Name: "Aux Data Essence", Symbol: "AuxDataEssence", Definition: "Identifies an Auxiliary Data Essence element", DefiningDocument: "SMPTE ST 429-14", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.177f0e7f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.177f0e7f", Name: "Frame-wrapped DMCVT Element", Symbol: "FrameWrappedDMCVTElement", Definition: "Identifies Frame-wrapped DMCVT Element", DefiningDocument: "SMPTE ST 2094-2", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.187f017f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.187f017f", Name: "DV-DIF Frame-wrapped", Symbol: "DVDIFFrameWrapped", Definition: "Identifies a frame-wrapped DV-DIF compound element", DefiningDocument: "SMPTE ST 383", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0d010301.187f027f": {UL: "urn:smpte:ul:060e2b34.01020101.0d010301.187f027f", Name: "DV-DIF Clip-wrapped", Symbol: "DVDIFClipWrapped", Definition: "Identifies a clip-wrapped DV-DIF compound element", DefiningDocument: "SMPTE ST 383", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020105.0e090502.017f017f": {UL: "urn:smpte:ul:060e2b34.01020105.0e090502.017f017f", Name: "Frame Wrapped ISXD Data", Symbol: "FrameWrappedISXDData", Definition: "Identifies Frame Wrapped ISXD Data Essence", DefiningDocument: "RDD 47", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0f020101.017f007f": {UL: "urn:smpte:ul:060e2b34.01020101.0f020101.017f007f", Name: "Frame Wrapped Binary Data", Symbol: "FrameWrappedBinaryData", Definition: "Identifies Frame Wrapped Binary Data", DefiningDocument: "MRX TestSpec", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.0101010c.0d01050d.01000000": {UL: "urn:smpte:ul:060e2b34.0101010c.0d01050d.01000000", Name: "Clip Wrapped Binary Data", Symbol: "ClipWrappedBinaryData", Definition: "Identifies Clip Wrapped Binary Data", DefiningDocument: "MRX TestSpec", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.0101010c.0d01050d.00000000": {UL: "urn:smpte:ul:060e2b34.0101010c.0d01050d.00000000", Name: "Clip Wrapped Text Data", Symbol: "ClipWrappedTextData", Definition: "Identifies Clip Wrapped Text Data", DefiningDocument: "MRX TestSpec", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.02050101.0d010301.04010100": {UL: "urn:smpte:ul:060e2b34.02050101.0d010301.04010100", Name: "System Metadata pack", Symbol: "SystemMetadatapack", Definition: "Identifies Clip Wrapped Text Data", DefiningDocument: "SMPTE 385", IsDeprecated: false},
	"urn:smpte:ul:060e2b34.01020101.0f020101.05000000": {UL: "urn:smpte:ul:060e2b34.01020101.0f020101.05000000", Name: "Metarex Manifest", Symbol: "MetarexManifest", Definition: "Identifies the Metarex Manifest", DefiningDocument: "MRX TestSpec", IsDeprecated: false},
}
