import api from "./axios";

export const getBoard = async (boardId: string) => {
    const response = await api.get(`/board/${boardId}`);

    return response.data;
};

export const getAllBoards = async (projectId: string) => {
    const response = await api.get(`/board/all-boards/${projectId}`);
    return response.data;
}