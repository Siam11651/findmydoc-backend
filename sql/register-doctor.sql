create or replace function register_doctor(param_id text)
returns void
language plpgsql
as $$
	begin
		update users
		set is_doctor = true
		where id = param_id;
	end;
$$;