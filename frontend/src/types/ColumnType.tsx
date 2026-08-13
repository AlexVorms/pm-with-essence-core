import type {TaskType} from "./TaskData.ts";


export interface ColumnType {
    id: string;
    name: string;
    isFinal: boolean;
    order: number;
    tasks: TaskType[];
}
