import api from "./axios";

export const getBoard = async (id: string) => {
    const response = await api.get(`/boards/${id}`);
    return response.data;
};