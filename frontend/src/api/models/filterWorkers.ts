export interface FilterWorkers {
    name: string;
    surname: string;
    patronymic: string;
    email: string;
    phone_number: string;
    orders_count_min: number;
    orders_count_max: number;
    created_at_begin: string;
    created_at_end: string;
}