import type {ColumnType} from "./ColumnType.tsx";

export interface BoardType {
    id: string;
    name: string;
    columns: ColumnType[];
}

export interface BoardPreview {
    id: string;
    name: string;
}