import api from "./axios.ts";

interface CreateTaskDto{
    name: string;
    description: string;
}
export const CreateTask = async (
    columnId: string,
    data: CreateTaskDto) => {

    const responce = await api.post(
    `/task/${columnId}`, data);

    return responce.data;

}