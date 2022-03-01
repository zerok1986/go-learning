import { Todo } from "../../entities/Todo"

type Props = {
    todos: Todo[];
}

export const TodoList: React.FC<Props> = ({ todos }) => {
    return (
        <ul>
            {
                todos.map((todo, idx) => (
                    <li key={idx}>{todo.title}</li>
                ))
            }
        </ul>
    )

}