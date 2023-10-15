export interface User {
	id: string;
	display_name: string;
	images: UserImage[];
	href: string;
	email: string;
	country: string;
}

interface UserImage {
	url: string;
	height: number;
	width: number;
}
