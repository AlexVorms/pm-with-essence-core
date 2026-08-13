import styles from './board.module.css'
import type {ColumnType} from "../../types/ColumnType.tsx";
import {Column} from "../column/Column.tsx";
import type {BoardType} from "../../types/BoardType.ts";
import { CirclePlus } from 'lucide-react';

interface BoardProps {
    columns: ColumnType[];
    board: BoardType;
}
export const Board = ({columns, board} : BoardProps) => {
    return (
        <div className={styles.page}>
            <div className={styles.header}>
            <h1>{board.name}</h1>
            </div>
            <div className={styles.board}>
                {columns.map((column) => (
                    <Column
                        key={column.id}
                        column={column}
                    />
                ))}
                <button className={styles.addColumnButton}>
                    <CirclePlus />
                    <span>Добавить колонку </span>
                </button>
            </div>
        </div>
    );
};