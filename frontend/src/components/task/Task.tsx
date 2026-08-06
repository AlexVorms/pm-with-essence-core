import type {TaskType} from "../../types/TaskData.ts";
import styles from "./task.module.css"
interface TaskProps {
    task: TaskType;
}
export const Task = ({task}: TaskProps) => {
    return(
        <div className={styles.task}>
            <h3 className={styles.title}>
                {task.name}
            </h3>

            {task.description && (
                <p className={styles.description}>
                    {task.description}
                </p>
            )}
        </div>
    );
};