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
