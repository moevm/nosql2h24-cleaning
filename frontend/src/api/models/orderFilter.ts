interface OrderFilter {
  user_id: string;
  address_city: string;
  address_street: string;
  address_building: string;
  address_entrance: string;
  address_floor: string;
  address_door_number: string;
  price_min: number;
  price_max: number;
  pollution_min: number;
  pollution_max: number;
  area_min: number;
  area_max: number;
  number_of_rooms_min: number;
  number_of_rooms_max: number;
  number_of_bathrooms_min: number;
  number_of_bathrooms_max: number;
  required_workers_min: number;
  required_workers_max: number;

  worker_name: string;
  worker_surname: string;
  worker_patronymic: string;
  statuses?: string[];
  services?: string[];
  date_time_begin: string;
  date_time_end: string;
}

export default OrderFilter;