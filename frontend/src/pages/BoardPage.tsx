import { useBoard } from "../hooks/useBoard";
import { useBoards } from "../hooks/useBoards";
import {Board} from "../components/board/Board.tsx";
import {BoardTabs} from "../components/board/BoardTabs/BoardTabs.tsx";


interface BoardPageProps {
    boardId: string;
    projectId: string;
}

export const BoardPage = ({ boardId, projectId }: BoardPageProps) => {

    const {
        data: boards,
        isLoading: isBoardsLoading,
    } = useBoards(projectId!);
    console.log(boards);
    const {
        data: board,
        isLoading: isBoardLoading,
        isError,
    } = useBoard(boardId);


    if (isBoardsLoading || isBoardLoading) {
        return <div>Загрузка...</div>;
    }

    if (isError || !board) {
        return <div>Не удалось загрузить доску</div>;
    }

    return (
            <div>
                {boards && (
                    <BoardTabs
                        boards={boards}
                        activeBoardId={boardId}
                        projectId={projectId}
                    />
                )}

                <Board columns = {board.columns} board = {board}></Board>
            </div>


    );
};