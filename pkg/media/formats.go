package media

import "github.com/photoprism/photoprism/pkg/fs"

// Formats maps file formats to general media types.
var Formats = map[fs.Type]Type{
	fs.ImageRaw:        Raw,
	fs.ImageDNG:        Raw,
	fs.ImageJPEG:       Image,
	fs.ImageJPEGXL:     Image,
	fs.ImageThumb:      Image,
	fs.ImagePNG:        Image,
	fs.ImageGIF:        Image,
	fs.ImageTIFF:       Image,
	fs.ImagePSD:        Image,
	fs.ImageBMP:        Image,
	fs.ImageMPO:        Image,
	fs.ImageAVIF:       Image,
	fs.ImageAVIFS:      Image,
	fs.ImageHEIF:       Image,
	fs.ImageHEIC:       Image,
	fs.ImageHEICS:      Image,
	fs.ImageWebP:       Image,
	fs.SidecarXMP:      Sidecar,
	fs.SidecarXML:      Sidecar,
	fs.SidecarAAE:      Sidecar,
	fs.SidecarYAML:     Sidecar,
	fs.SidecarJSON:     Sidecar,
	fs.SidecarText:     Sidecar,
	fs.SidecarInfo:     Sidecar,
	fs.SidecarMarkdown: Sidecar,
	fs.VectorSVG:       Vector,
	fs.VectorAI:        Vector,
	fs.VectorPS:        Vector,
	fs.VectorEPS:       Vector,
	fs.VideoWebM:       Video,
	fs.VideoAVC:        Video,
	fs.VideoHEVC:       Video,
	fs.VideoVVC:        Video,
	fs.VideoEVC:        Video,
	fs.VideoAVI:        Video,
	fs.VideoAV1:        Video,
	fs.VideoMPG:        Video,
	fs.VideoMJPG:       Video,
	fs.VideoMP2:        Video,
	fs.VideoMP4:        Video,
	fs.VideoM4V:        Video,
	fs.VideoMKV:        Video,
	fs.VideoMOV:        Video,
	fs.VideoMXF:        Video,
	fs.Video3GP:        Video,
	fs.Video3G2:        Video,
	fs.VideoFlash:      Video,
	fs.VideoAVCHD:      Video,
	fs.VideoBDAV:       Video,
	fs.VideoOGV:        Video,
	fs.VideoASF:        Video,
	fs.VideoWMV:        Video,
	fs.VideoDV:         Video,
	fs.TypeUnknown:     Sidecar,
}