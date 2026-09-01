import { useState } from "react";
import styles from "./addTaskModal.module.css"
interface CreateTaskModalProps {
    onClose: () => void;
    onSubmit: (name: string, description: string) => void;
}

export const AddTaskModal = ({
                                    onClose,
                                    onSubmit,
                                }: CreateTaskModalProps) => {
    const [name, setName] = useState("");
    const [description, setDescription] = useState("");

    const handleSubmit = (event: React.FormEvent) => {
        event.preventDefault();

        if (!name.trim()) {
            return;
        }

        onSubmit(name, description);
    };

    return (
        <div className={styles.modalOverlay} onClick={onClose}>
            <div
                className={styles.modal}
                onClick={(event) => event.stopPropagation()}
            >
                <div className={styles.modalHeader}>
                    <h2>Новая задача</h2>

                    <button
                        type="button"
                        onClick={onClose}
                    >
                        ×
                    </button>
                </div>

                <form onSubmit={handleSubmit}>
                    <label>
                        Название
                        <input
                            value={name}
                            onChange={(event) =>
                                setName(event.target.value)
                            }
                            placeholder="Введите название задачи"
                        />
                    </label>

                    <label>
                        Описание
                        <textarea
                            value={description}
                            onChange={(event) =>
                                setDescription(event.target.value)
                            }
                            placeholder="Введите описание"
                        />
                    </label>

                    <div className={styles.modalActions}>
                        <button
                            type="button"
                            onClick={onClose}
                        >
                            Отмена
                        </button>

                        <button type="submit">
                            Создать
                        </button>
                    </div>
                </form>
            </div>
        </div>
    );
};