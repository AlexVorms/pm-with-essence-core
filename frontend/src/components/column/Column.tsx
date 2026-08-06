import styles from "./column.module.css"
import type {ColumnType} from "../../types/ColumnData.tsx";
import { Task } from "../task/Task.tsx";

interface ColumnProps {
    column: ColumnType
}
export const Column = ({column}: ColumnProps) => {
    return(
    <div className={styles.column}>
        <div className={styles.columnTitle}>
            <h5>{column.name}</h5>
        </div>

        <div className={styles.columnBody}>
            {column.tasks.map((task) => (
                <Task
                    key={task.id}
                    task={task}
                />
            ))}
            <div>+ Добавить задачу</div>
        </div>
    </div>
    );
};