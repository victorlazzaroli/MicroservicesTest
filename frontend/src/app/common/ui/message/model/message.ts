export interface MessageI {
    id?: number;
    title: string;
    text: string;
    author: string;
    createdAt: string;
    updatedAt: string;
    likes: number;
    dislikes: number;
}