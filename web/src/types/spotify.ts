export interface ImageURL {
	url: string;
	height: number;
	width: number;
}

export interface Playlists {
	items: PlaylistItem[];
	limit: number;
	next: number;
	offset: number;
	previous: number;
	total: number;
}

export interface PlaylistItem {
	id: string;
	images: ImageURL[];
	name: string;
}

export interface PlaylistTracks {
	items: PlaylistItem[];
	limit: number;
	offset: number;
	total: number;
}

export interface PlaylistItem {
	track: Track;
}

export interface Track {
	id: string;
	name: string;
	duration_ms: number;
	artists: Artist[];
}

export interface Artist {
	id: string;
	name: string;
}

export interface DownloadLink {
	link: string;
}
