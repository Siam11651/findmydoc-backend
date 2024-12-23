create or replace function update_location(param_id text, param_latitude float8, param_longitude float8)
returns void
language plpgsql
as $$
	begin
		update users
		set "last_point" = point(param_latitude, param_longitude), "last_time" = now()
		where "id" = param_id;
	end;
$$;