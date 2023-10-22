import { ImageURL } from './spotify';

export interface User {
	id: string;
	display_name: string;
	images: ImageURL[];
	href: string;
	email: string;
	country: string;
}
