create or replace function is_doctor(param_id text)
returns bool
language plpgsql
as $$
	declare
		_result bool := null;
	begin
		select is_doctor into _result
		from users
		where id = param_id;

		return _result;		
	end;
$$;