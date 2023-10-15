export interface AccessTokens {
	accessToken: string;
	expiresIn: number;
	refreshToken: string;
}

export interface RefreshAccessTokens {
	access_token: string;
	expires_in: number;
}
