export interface MessageI {
    id?: number;
    title: string;
    text: string;
    author: string;
    author_id: number;
    createdAt: string;
    updatedAt: string;
    likes: number;
    dislikes: number;
}