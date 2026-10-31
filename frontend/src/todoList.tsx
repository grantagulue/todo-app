import { useState } from "react";
import { useQuery } from "@tanstack/react-query";
import { apiClient } from "../api/requestProcessor";

interface Todo {
	id: number;
	title: string;
	completed: boolean;
}

const TodoList = () => {
	const [tasks, setTasks] = useState<string[]>([
		"Walk to the Store",
		"Take out the trash",
	]);

	const {
		data: todos,
		isLoading,
		isError,
	} = useQuery({
		queryKey: ["todos"],
		queryFn: () => apiClient.get("/todos").then((response) => response.data),
	});

	const addTodo = async (title: string) => {
		try {
			await apiClient.post("/todos", { title });
		} catch (error) {
			console.error("Error adding todo:", error);
		}
	};

	const [newTask, setNewTask] = useState<string>("");
	return (
		<>
			<div>
				<div>
					<input
						type="text"
						value={newTask}
						onChange={(e) => {
							setNewTask(e.target.value);
						}}
					/>
					<button
						onClick={() => {
							setTasks([...tasks, newTask]);
							setNewTask("");

							addTodo(newTask);
						}}
					>
						Add
					</button>
				</div>
				{todos?.map((task: Todo, index: number) => {
					return <li key={index}>{task.title}</li>;
				})}
				: {isLoading && <p>Loading...</p>} :{" "}
				{isError && <p>Error fetching todos</p>}
			</div>
		</>
	);
};

export default TodoList;
