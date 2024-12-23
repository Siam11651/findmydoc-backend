create or replace function register(param_id text)
returns void
language plpgsql
as $$
	declare
		exist_id text := null;
	begin
		select id into exist_id
		from users
		where id = param_id;

		if exist_id is null
		then
			insert into users(id)
			values (param_id);
		end if;
	end;
$$;