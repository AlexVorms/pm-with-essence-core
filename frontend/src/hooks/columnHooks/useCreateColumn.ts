import {useMutation, useQueryClient} from "@tanstack/react-query";
import {CreateColumn} from "../../api/columnAPI.ts";

export const useCreateColumn = () => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({
            boardId,
            name,
        }: {
            boardId: string;
            name: string;
        }) => CreateColumn(boardId, {name}),
        onSuccess: (_, variables) => {
            queryClient.invalidateQueries({
                queryKey:["board", variables.boardId],
            });
        },
    });
};