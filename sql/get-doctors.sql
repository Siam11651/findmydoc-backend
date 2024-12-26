drop function get_doctors;

create or replace function get_doctors(param_id text, param_point point)
returns table (
	id text,
	latitude float8,
	longitude float8
)
language plpgsql
as $$
	begin	
		return query select users.id, users.last_point[0], users.last_point[1] 
		from users
		where 
			users.id <> param_id
			and users.is_doctor = true
			and get_distance(param_point, users.last_point) <= 3.0
			and EXTRACT(SECOND FROM (now() - users.last_time)) <= 30;
	end;
$$;