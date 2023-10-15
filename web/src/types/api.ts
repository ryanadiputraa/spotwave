export interface SuccessResponse<T> {
	message: string;
	data: T;
}

export interface ErrorResponse {
	err_code: string;
	message: string;
}
