import styles from './board.module.css'
import type {ColumnType} from "../../types/ColumnType.tsx";
import {Column} from "../column/Column.tsx";
import type {BoardType} from "../../types/BoardType.ts";
import { CirclePlus } from 'lucide-react';
import {useState} from "react";
import {AddColumnButton} from "../column/AddColumnButton/AddColumnButton.tsx";
import {useCreateColumn} from "../../hooks/columnHooks/useCreateColumn.ts";

interface BoardProps {
    columns: ColumnType[];
    board: BoardType;
}
export const Board = ({columns, board} : BoardProps) => {
   const [isAddingColumn, setIsAddingColumn] = useState(false);

   const {mutate: createColumn} = useCreateColumn();
    const handleCancelAdding = () => {
        setIsAddingColumn(false);
    }
    const handleCreateColumn = (name: string) => {
        console.log(name);
        createColumn(
            {
            boardId: board.id,
            name: name,
        },
            {
                onSuccess: () => {
                    setIsAddingColumn(false);
                }
            }
        )
    }
    return (
        <div className={styles.page}>

            <div className={styles.board}>
                {columns.map((column) => (
                    <Column
                        key={column.id}
                        column={column}
                    />
                ))}
                {isAddingColumn ?
                    <AddColumnButton
                    onCancel={handleCancelAdding}
                    onSubmit={handleCreateColumn} />
                    : (
                    <button className={styles.addColumnButton}
                            onClick={() => setIsAddingColumn(true)}>
                        <CirclePlus />
                        <span>Добавить колонку </span>
                    </button>
                )}

            </div>
        </div>
    );
};