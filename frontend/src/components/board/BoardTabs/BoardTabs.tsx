import type {BoardPreview} from "../../../types/BoardType.ts";
import styles from "./boardTabs.module.css"
import { useNavigate } from "react-router-dom";

interface BoardTabsProps {
    boards: BoardPreview[]
    activeBoardId: string
    projectId: string
}
export const BoardTabs = ({
                              boards,
                              activeBoardId,
                              projectId}: BoardTabsProps) => {

    const navigate = useNavigate();
    const handleBoardClick = (id: string) => {
        navigate(`/projects/${projectId}/boards/${id}`);
    };

    return (
        <div className={styles.tabs}>
            {boards.map(board => (
                <button key={board.id}
                        className={`${styles.tab} ${
                            board.id === activeBoardId
                                ? styles.active
                                : ""
                        }`}
                        onClick={() => handleBoardClick(board.id)}
                >
                    {board.name}
                </button>
            ))}
        </div>
    );
};