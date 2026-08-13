import type {BoardPreview} from "../../../types/BoardType.ts";
import styles from "./boardTabs.module.css"
interface BoardTabsProps {
    boards: BoardPreview[]
}
export const BoardTabs = ({boards}: BoardTabsProps) => {
    return (
        <div className={styles.tabs}>
            {boards.map(board => (
                <button key={board.id}
                className={styles.tab}>
                    {board.name}
                </button>
            ))}
        </div>
    );
};