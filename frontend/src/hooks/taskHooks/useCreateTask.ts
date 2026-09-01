import {useMutation, useQueryClient} from "@tanstack/react-query";
import {CreateTask} from "../../api/taskAPI.ts";

export const useCreateTask = () => {
    const queryClient = useQueryClient();
    return useMutation({
        mutationFn: ({
            columnId,
            name,
            description
        }: {
            columnId: string;
            name: string;
            description: string;
        }) => CreateTask(columnId,{
            name:name,
            description:description,
        }),
        onSuccess:(_, variables) => {
            queryClient.invalidateQueries({
                queryKey:
                }
            )
        }
    })
}