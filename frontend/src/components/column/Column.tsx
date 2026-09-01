import styles from "./column.module.css"
import type {ColumnType} from "../../types/ColumnType.tsx";
import { Task } from "../task/Task.tsx";
import {useState} from "react";
import {AddTaskModal} from "../task/AddTaskModal/AddTaskModal.tsx";
interface ColumnProps {
    column: ColumnType
}
export const Column = ({column}: ColumnProps) => {
   const  [isAddingTask, setIsAddingTask] = useState(false);
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
                <button onClick={() => setIsAddingTask(true)}
                    className={styles.addTaskButton}>
                    + Добавить задачу
                </button>
        </div>
        {isAddingTask && (<AddTaskModal
            onClose={() =>  setIsAddingTask(false)}
            onSubmit={() => {}}/>)}
    </div>
    );
};