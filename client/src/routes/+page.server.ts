import { generateUUID } from '$lib/genereateUid';
import { addTodo, getTodos, updateTodo } from '$lib/server/Api';
import type { Actions } from './$types';
import type { Todo } from './interface';


export function load() {
  const data = getTodos();
  return {
    todos: data
  }
}


export const actions = {
  addTodo: async ( {request} ) => {
    const data = await request.formData();

    console.log(data.get('title'));

    // create a todo
    const todo: Todo = {
      id: generateUUID(),
      title: data.get('title')?.toString(),
      description: data.get('description')?.toString(),
      completed: false
    }


    // send the data to the server
    const res = addTodo(todo);

    // return the response
    return res
  },
  updateTodo: async ( {request} ) => {
    const data = await request.formData();
    console.log(data);
    // create a todo
    const todo: Todo = {
      id: data.get('id')?.toString(),
      title: data.get('title')?.toString(),
      description: data.get('description')?.toString(),
      completed: true
    }

    const res = await updateTodo(todo.id, todo);

    return res

  },
  deleteTodo: async ( {request} ) => {
    const data = await request.formData();
    console.log(data);

  }
} 