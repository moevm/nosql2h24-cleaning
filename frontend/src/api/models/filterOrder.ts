export interface FilterOrder {
  user_name: string;
  user_surname: string;
  user_patronymic: string;
  address_city: string;
  address_street: string;
  address_building: string;
  address_entrance: string;
  address_floor: string;
  address_door_number: string;
  worker_name: string;
  worker_surname: string;
  worker_patronymic: string;
  statuses?: string[];
  services?: string[];
  date_time_begin?: string;
  date_time_end?: string;
}