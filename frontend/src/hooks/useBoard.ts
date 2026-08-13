import {useQuery} from "@tanstack/react-query";
import {getBoard} from "../api/boardAPI.ts";


export const useBoard = (boardId: string) => {
    return useQuery({
        queryKey:["board", boardId],
        queryFn: () => getBoard(boardId),
    })
}