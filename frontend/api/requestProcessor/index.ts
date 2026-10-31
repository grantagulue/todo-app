import axios, { type AxiosInstance } from "axios";

declare global {
	interface ImportMetaEnv {
		readonly VITE_BASE_URL: string;
	}

	interface ImportMeta {
		readonly env: ImportMetaEnv;
	}
}

const apiUrl: string = import.meta.env.VITE_BASE_URL;

export const apiClient: AxiosInstance = axios.create({
	baseURL: apiUrl,
});

apiClient.interceptors.request.use(
	(config) => {
		const token = localStorage.getItem("accessToken");

		if (token) {
			config.headers.Authorization = `Bearer ${token}`;
		}
		return config;
	},
	(error) => {
		return Promise.reject(error);
	},
);

apiClient.interceptors.response.use(
	(response) => {
		return response.data;
	},
	async (error) => {
		const originalRequest = error.config;

		if (error.response?.status === 401 && !originalRequest._retry) {
			originalRequest._retry = true;

			try {
				const { newAccessToken }: { newAccessToken: string } = await axios.post(
					`${apiUrl}/auth/refresh`,
				);
				localStorage.setItem("accessToken", newAccessToken);

				originalRequest.headers.Authorization = `Bearer ${newAccessToken}`;
				return apiClient(originalRequest);
			} catch (refreshError) {
				localStorage.removeItem("accessToken");
				window.location.href = "/login";
				return Promise.reject(refreshError);
			}
		}

		return Promise.reject(error);
	},
);
