interface OrderFilter {
  user_id: string;
  address_city: string;
  address_street: string;
  address_building: string;
  address_entrance: string;
  address_floor: string;
  address_door_number: string;
  worker_name: string;
  worker_surname: string;
  worker_patronymic: string;
  statuses: string[];
  date_time: Date;
}

export default OrderFilter;