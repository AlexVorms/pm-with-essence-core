import {useQuery} from "@tanstack/react-query";
import {getAllBoards} from "../api/boardAPI.ts";

export const useBoards = (projectId:string) => {
    return useQuery({
        queryKey: ["boards", projectId],
        queryFn: () => getAllBoards(projectId),
    })
}