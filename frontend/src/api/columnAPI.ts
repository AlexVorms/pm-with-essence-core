import api from "./axios.ts";
interface CreateColumnDto {
    name: string;
}
export const CreateColumn = async (
    boardId: string,
    data: CreateColumnDto
) => {
    const response = await api.post(
        `/column/${boardId}`, data);
    return response.data;
}

