import style from "./addColumnButton.module.css"
import {useState} from "react";
interface AddColumnButtonProps {
    onCancel: () => void;
    onSubmit: (name: string) => void;
}
export const AddColumnButton = ({
    onCancel,
    onSubmit,
}: AddColumnButtonProps) => {
    const [columnName, setColumnName] = useState("");

    const handleSubmit = () =>{
        console.log(columnName)
        onSubmit(columnName);
    }
    return (
        <div className={style.AddColumnButton}>
            <input
                value={columnName}
                placeholder="Название колонки"
            onChange={(e) => setColumnName(e.target.value)}/>

            <div >
                <button
                    className={style.AddButton}
                onClick={handleSubmit}>
                    Добавить
                </button>

                <button className={style.CancelButton} onClick={onCancel}>
                    ×
                </button>
            </div>
        </div>
    );
};